package main

import (
    "encoding/json"
    "fmt"
    "math"
    "strconv"
    "strings"
)

const EarthRadiusKm = 6371.0
const fullFrameWidthMM = 36.0 // assume 35mm-equivalent focal length

// ---- Generic helpers ----

// Convert various types to float64
func toFloat64(v interface{}) float64 {
    switch val := v.(type) {
    case float64:
        return val
    case float32:
        return float64(val)
    case int:
        return float64(val)
    case int64:
        return float64(val)
    case string:
        f, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
        return f
    case json.Number:
        f, _ := val.Float64()
        return f
    default:
        return 0
    }
}

// Return fallback if string is empty
func ifEmpty(s string, fallback string) string {
    if strings.TrimSpace(s) == "" {
        return fallback
    }
    return s
}

// ---- Angle/length helpers ----

func degreesToRadians(deg float64) float64 {
    return deg * math.Pi / 180.0
}

func radiansToDegrees(rad float64) float64 {
    return rad * 180.0 / math.Pi
}

// Horizontal FOV from 35mm-equivalent focal length (mm), using 36 mm full-frame width.
func hfovDegreesFromFocalEq(focalEqMM float64) float64 {
    if focalEqMM <= 0 {
        return 50.0 // safe default angle
    }
    return 2.0 * radiansToDegrees(math.Atan((fullFrameWidthMM / 2.0) / focalEqMM))
}

// Compute a short, realistic cone length based on the heading line distance.
// Rule: cone length = heading line length / 7. Fallback to 120 m if line distance is missing.
func coneLengthFromHeading(destDistanceM float64) float64 {
    if destDistanceM <= 0 {
        return 120.0
    }
    return destDistanceM / 7.0
}

// ---- Map geometry helpers ----

// Great-circle forward projection (from lat/lon, bearing deg, distance meters) -> dest lat/lon
func computeOffsetPoint(lat, lon, bearingDeg, distanceM float64) [2]float64 {
    bearingRad := degreesToRadians(bearingDeg)
    latRad := degreesToRadians(lat)
    lonRad := degreesToRadians(lon)
    distanceKm := distanceM / 1000.0

    destLatRad := math.Asin(math.Sin(latRad)*math.Cos(distanceKm/EarthRadiusKm) +
        math.Cos(latRad)*math.Sin(distanceKm/EarthRadiusKm)*math.Cos(bearingRad))

    destLonRad := lonRad + math.Atan2(
        math.Sin(bearingRad)*math.Sin(distanceKm/EarthRadiusKm)*math.Cos(latRad),
        math.Cos(distanceKm/EarthRadiusKm)-math.Sin(latRad)*math.Sin(destLatRad),
    )

    return [2]float64{radiansToDegrees(destLatRad), radiansToDegrees(destLonRad)}
}

// ---- Heading line projection ----

// Compute heading line from FinalYaw. Sets DestLat/DestLon/DestBearing/DestDistanceM/DestDistanceStr.
func computeProjection(f FlightData, cfg UserConfig) FlightData {
    if f.Lat == 0 || f.Lon == 0 {
        return f
    }
    // Distance in meters based on config (Distance + Unit)
    var distM float64
    if cfg.Unit == "m" {
        distM = cfg.Distance
    } else {
        distM = cfg.Distance * 1000.0
    }
    if distM <= 0 {
        distM = 5000.0
    }

    // Bearing preference: FinalYaw, else FlightYaw, else GimbalYaw
    bearing := f.FinalYaw
    if bearing == 0 && f.FlightYaw != 0 {
        bearing = f.FlightYaw
    }
    if bearing == 0 && f.GimbalYaw != 0 {
        bearing = f.GimbalYaw
    }
    if bearing == 0 {
        return f
    }

    dest := computeOffsetPoint(f.Lat, f.Lon, bearing, distM)
    f.DestLat = dest[0]
    f.DestLon = dest[1]
    f.DestBearing = bearing
    f.DestDistanceM = distM
    f.DestDistanceStr = fmt.Sprintf("%.1f %s", distM, cfg.Unit)
    return f
}

// ---- Cones (POV wedges) ----

// computeOpticalConePolygon builds a realistic POV cone using focal length (HFOV),
// altitude, and gimbal pitch. It decides whether the cone starts at the pin or
// at the ground intercept depending on pitch angle.
// computeOpticalConePolygon builds a realistic POV cone using focal length (HFOV),
// altitude, and gimbal pitch. It respects DJI's convention: negative pitch = down,
// near 0 = horizon, positive = up (sky).
func computeOpticalConePolygon(f FlightData, fallbackSpreadDeg float64) [][2]float64 {
    if f.Lat == 0 || f.Lon == 0 {
        return nil
    }

    // Bearing preference
    yaw := f.FinalYaw
    if yaw == 0 && f.FlightYaw != 0 { yaw = f.FlightYaw }
    if yaw == 0 && f.GimbalYaw != 0 { yaw = f.GimbalYaw }
    if yaw == 0 { return nil }

    // Horizontal FOV from focal length
    var hfov float64
    if f.FocalLengthMM > 0 {
        hfov = hfovDegreesFromFocalEq(f.FocalLengthMM)
    } else {
        hfov = fallbackSpreadDeg
    }
    if hfov <= 1 {
        hfov = 30
    }

    h := f.Altitude
    pitch := f.GimbalPitchDeg

    var startAtPin bool
    var centerDist float64

    switch {
    case pitch < -5: // camera pitched down
        startAtPin = true
        t := degreesToRadians(-pitch) // use absolute downward angle
        if t > 0 {
            centerDist = h / math.Tan(t)
        }
        if centerDist < 20 { centerDist = 20 }
        if centerDist > 1000 { centerDist = 1000 }

    case math.Abs(pitch) <= 5: // near horizon
        startAtPin = false
        centerDist = h * 20
        if centerDist < 100 { centerDist = 100 }
        if centerDist > 2000 { centerDist = 2000 }

    case pitch > 5: // camera tilted up, filming sky
        return nil

    default:
        startAtPin = true
        centerDist = 200
    }

    // Build arc
    half := hfov / 2.0
    samples := 25
    step := hfov / float64(samples-1)

    arc := make([][2]float64, 0, samples)
    for i := 0; i < samples; i++ {
        bearing := yaw - half + float64(i)*step
        dest := computeOffsetPoint(f.Lat, f.Lon, bearing, centerDist)
        arc = append(arc, [2]float64{dest[0], dest[1]})
    }

    if startAtPin {
        pts := make([][2]float64, 0, 1+len(arc))
        pts = append(pts, [2]float64{f.Lat, f.Lon})
        pts = append(pts, arc...)
        return pts
    }
    return arc
}




// Legacy wedge: short wedge centered at FinalYaw with a fixed spread angle.
func computeConePolygon(f FlightData, spreadDeg float64) [][2]float64 {
    if f.Lat == 0 || f.Lon == 0 {
        return nil
    }
    yaw := f.FinalYaw
    if yaw == 0 && f.FlightYaw != 0 {
        yaw = f.FlightYaw
    }
    if yaw == 0 && f.GimbalYaw != 0 {
        yaw = f.GimbalYaw
    }
    if yaw == 0 {
        return nil
    }

    coneLengthM := coneLengthFromHeading(f.DestDistanceM)
    if coneLengthM <= 0 {
        coneLengthM = 120.0
    }

    var points [][2]float64
    halfSpread := spreadDeg / 2
    steps := 5
    stepSize := spreadDeg / float64(steps-1)

    for i := 0; i < steps; i++ {
        angle := yaw - halfSpread + float64(i)*stepSize
        dest := computeOffsetPoint(f.Lat, f.Lon, angle, coneLengthM)
        points = append(points, [2]float64{dest[0], dest[1]})
    }
    return points
}
