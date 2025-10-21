package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/fs"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
    "strconv"
    "strings"
    "time"
)

type FlightData struct {
    FileName        string
    AbsolutePath    string
    Time            string
    Lat             float64
    Lon             float64
    Altitude        float64
    FlightYaw       float64
    GimbalYaw       float64
    FinalYaw        float64
    DestLat         float64
    DestLon         float64
    DestBearing     float64
    DestDistanceM   float64
    DestDistanceStr string
    HasTelemetry    bool
    ConePoints      [][2]float64
    FocalLengthMM   float64
    CameraModel     string
    GimbalPitchDeg  float64
}

type SRTRecord struct {
    Lat     float64
    Lon     float64
    Alt     float64
    Heading float64
    TimeISO string
    Focal   float64
    Pitch   float64
}

func processFiles(dir string, cfg UserConfig) ([]FlightData, error) {
    var flights []FlightData

    err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if d.IsDir() {
            if path != dir && !cfg.Recurse {
                return fs.SkipDir
            }
            return nil
        }

        extOrig := filepath.Ext(path)           // preserve original case (e.g., ".MP4")
        ext := strings.ToLower(extOrig)         // detect type case-insensitively
        base := strings.TrimSuffix(path, extOrig)

        isImage := ext == ".jpg" || ext == ".jpeg" || ext == ".dng" || ext == ".png"
        isVideo := ext == ".mp4"

        if isImage && !cfg.ScanImages {
            return nil
        }
        if isVideo && !cfg.ScanVideos {
            return nil
        }

        fmt.Printf("%s %s\n", lang["processing_file"], path)

        switch {
        case isImage:
            f, err := processJPG(path, cfg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", path, err)
                return nil
            }
            flights = append(flights, f)

        case isVideo:
            srtPath := base + ".srt" // correct pairing: "DJI_..._D.srt"
            if _, err := os.Stat(srtPath); err == nil {
                f, err := flightDataFromSRT(path, srtPath, cfg)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Error reading SRT %s: %v\n", srtPath, err)
                    return nil
                }
                flights = append(flights, f)
            } else {
                fd := fallbackMP4Data(path)
                flights = append(flights, fd)
            }
        }
        return nil
    })

    return flights, err
}


func processJPG(filePath string, cfg UserConfig) (FlightData, error) {
    cmd := exec.Command(tempPath("exiftool.exe"),
        "-json", "-n",
        "-GPSLatitude", "-GPSLongitude", "-GPSAltitude",
        "-FlightYawDegree", "-GimbalYawDegree", "-DateTimeOriginal",
        "-FocalLength", "-Model", "-GimbalPitchDegree",
        filePath)

    output, err := cmd.Output()
    if err != nil {
        return FlightData{}, fmt.Errorf("ExifTool error: %v", err)
    }

    var metadataList []map[string]interface{}
    if err := json.Unmarshal(output, &metadataList); err != nil {
        return FlightData{}, fmt.Errorf("error parsing ExifTool JSON: %v", err)
    }
    if len(metadataList) == 0 {
        return FlightData{}, fmt.Errorf("no metadata found in %s", filePath)
    }

    md := metadataList[0]
    data := FlightData{
        FileName:       filepath.Base(filePath),
        AbsolutePath:   filePath,
        Lat:            toFloat64(md["GPSLatitude"]),
        Lon:            toFloat64(md["GPSLongitude"]),
        Altitude:       toFloat64(md["GPSAltitude"]),
        FlightYaw:      toFloat64(md["FlightYawDegree"]),
        GimbalYaw:      toFloat64(md["GimbalYawDegree"]),
        FinalYaw:       toFloat64(md["FlightYawDegree"]),
        HasTelemetry:   true,
        FocalLengthMM:  toFloat64(md["FocalLength"]),
        GimbalPitchDeg: toFloat64(md["GimbalPitchDegree"]),
    }
    if m, ok := md["Model"].(string); ok {
        data.CameraModel = m
    }

    if rawTime, ok := md["DateTimeOriginal"].(string); ok {
        if t, err := time.Parse("2006:01:02 15:04:05", rawTime); err == nil {
            data.Time = t.Format("2006-01-02 15:04:05")
        } else {
            data.Time = "N/A"
        }
    } else {
        data.Time = "N/A"
    }

    // Compute heading line projection
    data = computeProjection(data, cfg)
    // Compute realistic POV cone (altitude + pitch aware)
    data.ConePoints = computeOpticalConePolygon(data, 50)

    return data, nil
}

func flightDataFromSRT(mp4Path, srtPath string, cfg UserConfig) (FlightData, error) {
    rec, err := parseDJISRTFirstPoint(srtPath)
    if err != nil {
        return FlightData{}, err
    }

    fd := FlightData{
        FileName:       filepath.Base(mp4Path),
        AbsolutePath:   mp4Path,
        Time:           ifEmpty(rec.TimeISO, "N/A"),
        Lat:            rec.Lat,
        Lon:            rec.Lon,
        Altitude:       rec.Alt,
        FlightYaw:      rec.Heading,
        FinalYaw:       rec.Heading,
        DestBearing:    rec.Heading,
        HasTelemetry:   true,
        FocalLengthMM:  rec.Focal,
        GimbalPitchDeg: rec.Pitch,
    }

    // fallback: parse time from filename if missing
    if fd.Time == "N/A" || fd.Time == "" {
        base := filepath.Base(mp4Path)
        re := regexp.MustCompile(`(\d{14})`)
        if match := re.FindStringSubmatch(base); match != nil {
            if t, err := time.Parse("20060102150405", match[1]); err == nil {
                fd.Time = t.Format("2006-01-02 15:04:05")
            }
        }
    }

    fd = computeProjection(fd, cfg)
    fd.ConePoints = nil // videos don’t draw cones

    return fd, nil
}


func fallbackMP4Data(path string) FlightData {
    fd := FlightData{
        FileName:     filepath.Base(path),
        AbsolutePath: path,
        HasTelemetry: false,
        Time:         "N/A",
    }
    base := filepath.Base(path)
    re := regexp.MustCompile(`(\d{14})`) // e.g., 20250420174947
    if match := re.FindStringSubmatch(base); match != nil {
        if t, err := time.Parse("20060102150405", match[1]); err == nil {
            fd.Time = t.Format("2006-01-02 15:04:05")
        }
    }
    return fd
}

func parseDJISRTFirstPoint(srtPath string) (*SRTRecord, error) {
    f, err := os.Open(srtPath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    var rec SRTRecord

    // Case-insensitive, permissive regexes
    focalRe := regexp.MustCompile(`(?i)focal_len:\s*([-\d\.]+)`)
    latRe   := regexp.MustCompile(`(?i)latitude:\s*([-\d\.]+)`)
    lonRe   := regexp.MustCompile(`(?i)longitude:\s*([-\d\.]+)`)
    relAltRe:= regexp.MustCompile(`(?i)rel_alt:\s*([-\d\.]+)`)
    absAltRe:= regexp.MustCompile(`(?i)abs_alt:\s*([-\d\.]+)`)
    pitchRe := regexp.MustCompile(`(?i)gimbal_pitch:\s*([-\d\.]+)`)
    yawRe   := regexp.MustCompile(`(?i)(yaw|heading):\s*([-\d\.]+)`)

    timeLine := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}\s+\d{2}:\d{2}:\d{2}\.\d+`)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }

        // Timestamp line (your SRT shows a raw ISO-like line before telemetry)
        if rec.TimeISO == "" && timeLine.MatchString(line) {
            rec.TimeISO = line
            continue
        }

        // Telemetry line (line 5 in your sample)
        if strings.Contains(strings.ToLower(line), "latitude:") {
            if m := focalRe.FindStringSubmatch(line); m != nil {
                rec.Focal, _ = strconv.ParseFloat(m[1], 64)
            }
            if m := latRe.FindStringSubmatch(line); m != nil {
                rec.Lat, _ = strconv.ParseFloat(m[1], 64)
            }
            if m := lonRe.FindStringSubmatch(line); m != nil {
                rec.Lon, _ = strconv.ParseFloat(m[1], 64)
            }
            if m := relAltRe.FindStringSubmatch(line); m != nil {
                rec.Alt, _ = strconv.ParseFloat(m[1], 64)
            } else if m := absAltRe.FindStringSubmatch(line); m != nil {
                rec.Alt, _ = strconv.ParseFloat(m[1], 64)
            }
            if m := pitchRe.FindStringSubmatch(line); m != nil {
                rec.Pitch, _ = strconv.ParseFloat(m[1], 64)
            }
            if m := yawRe.FindStringSubmatch(line); m != nil && len(m) >= 3 {
                rec.Heading, _ = strconv.ParseFloat(m[2], 64)
            }
            break
        }
    }

    if rec.Lat == 0 || rec.Lon == 0 {
        return nil, fmt.Errorf("no valid GPS coordinates in %s", srtPath)
    }
    return &rec, nil
}


