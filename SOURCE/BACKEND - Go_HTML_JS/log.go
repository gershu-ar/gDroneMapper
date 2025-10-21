package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func startLog(cfg UserConfig, flights []FlightData, errors int, htmlFileName string) error {
    if !cfg.EnableLogging {
        return nil
    }

    if err := os.MkdirAll(cfg.LogOutputFolder, 0755); err != nil {
        return fmt.Errorf("error creating log folder: %v", err)
    }

    logName := strings.Replace(htmlFileName, ".html", ".log", 1)
    logPath := filepath.Join(cfg.LogOutputFolder, logName)
    f, err := os.Create(logPath)
    if err != nil {
        return fmt.Errorf("error creating log file: %v", err)
    }
    defer f.Close()

    header := fmt.Sprintf(
        "%s %s\nInput Folder: %s\nOutput Folder: %s\nOutput File: %s\nLog Folder: %s\nFiles Processed: %d\nErrors: %d\n---\n",
        AppName, AppVersion,
        cfg.InputDirectory,
        cfg.OutputDirectory,
        htmlFileName,
        cfg.LogOutputFolder,
        len(flights),
        errors,
    )
    f.WriteString(header)

    for i, d := range flights {
        line := fmt.Sprintf("%02d, %s, %s, %.2f°, %.1fm, %.6f, %.6f\n",
            i+1, d.FileName, d.Time, d.FinalYaw, d.Altitude, d.Lat, d.Lon)
        f.WriteString(line)
    }

    return nil
}
