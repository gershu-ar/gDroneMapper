package main

import (
    "fmt"
    "path/filepath"
    "time"
)

func main() {
    // Load config from external file
    cfg, err := loadConfig(ConfigFileName)
    if err != nil {
        fmt.Println("Config error:", err)
        return
    }

    // Load language strings from embedded JSON
    if err := loadLanguageFromEmbed(cfg.Language); err != nil {
        fmt.Println("Language error:", err)
        return
    }

    // Deploy ExifTool assets to temp folder
    if err := deployExifAssetsIfMissing(); err != nil {
        fmt.Println("Deployment error:", err)
        return
    }

    fmt.Println(lang["processing_files"], cfg.InputDirectory)

    // Process files in input directory
    flights, err := processFiles(cfg.InputDirectory, cfg)
    if err != nil {
        fmt.Println(lang["error_processing"], err)
        return
    }
    if len(flights) == 0 {
        fmt.Println(lang["no_flights"])
        return
    }

    // Generate timestamp-based output name
    timestamp := time.Now().Format("20060102_150405")
    htmlName := fmt.Sprintf("gDroneMapper_%s.html", timestamp)

    // Generate map output
    if err := generateConcentratedMap(flights, cfg, lang, htmlName); err != nil {
        fmt.Println(lang["error_generating_map"], err)
        return
    }

    // Log results if enabled
    if err := startLog(cfg, flights, 0, htmlName); err != nil {
        fmt.Println("Log error:", err)
    }

    fmt.Println(lang["map_generated"], filepath.Join(cfg.OutputDirectory, htmlName))
    fmt.Println(lang["done"])
}
