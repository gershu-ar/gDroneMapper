package main

import (
    "encoding/json"
    "fmt"
    "os"
)

const AppName = "gDroneMapper"
const AppVersion = "Versión 1.0 Final"
const ConfigFileName = "config.json"

type UserConfig struct {
    InputDirectory   string  `json:"inputDirectory"`
    OutputDirectory  string  `json:"outputDirectory"`
    LogOutputFolder  string  `json:"log_output_folder"`
    Recurse          bool    `json:"recurse"`
    Distance         float64 `json:"distance"`
    Unit             string  `json:"unit"`
    DrawHeading      bool    `json:"drawHeading"`
    Language         string  `json:"language"`
    ScanImages       bool    `json:"scan_images"`
    ScanVideos       bool    `json:"scan_videos"`
    EnableLogging    bool    `json:"enable_logging"`
    ShowCones        bool    `json:"show_cones"`
}

func loadConfig(path string) (UserConfig, error) {
    defaultConfig := UserConfig{
        InputDirectory:  "./input",
        OutputDirectory: "./output",
        LogOutputFolder: "./log",
        Recurse:         true,
        Distance:        5.0,
        Unit:            "km",
        DrawHeading:     true,
        Language:        "en",
        ScanImages:      true,
        ScanVideos:      true,
        EnableLogging:   true,
        ShowCones:       true,
    }

    if _, err := os.Stat(path); os.IsNotExist(err) {
        data, _ := json.MarshalIndent(defaultConfig, "", "  ")
        os.WriteFile(path, data, 0644)
        fmt.Printf("Config file created: %s\n", path)
        return defaultConfig, nil
    }

    data, err := os.ReadFile(path)
    if err != nil {
        return defaultConfig, err
    }

    var cfg UserConfig
    if err := json.Unmarshal(data, &cfg); err != nil {
        return defaultConfig, err
    }

    if !cfg.ScanImages && !cfg.ScanVideos {
        return defaultConfig, fmt.Errorf("invalid config: both scan_images and scan_videos are false")
    }

    return cfg, nil
}
