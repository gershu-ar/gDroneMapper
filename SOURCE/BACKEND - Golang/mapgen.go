package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "html/template"
    "os"
    "path/filepath"
)

func generateConcentratedMap(flights []FlightData, cfg UserConfig, lang map[string]string, htmlFileName string) error {
    tmplBytes, err := embeddedAssets.ReadFile("template.html")
    if err != nil {
        return fmt.Errorf("error reading embedded template: %v", err)
    }

    tmpl, err := template.New("map").Parse(string(tmplBytes))
    if err != nil {
        return fmt.Errorf("error parsing HTML template: %v", err)
    }

    jsonData, err := json.Marshal(flights)
    if err != nil {
        return fmt.Errorf("error marshaling flight data: %v", err)
    }

    data := struct {
        Title           string
        AppName         string
        AppVersion      string
        GeneratedTime   string
        DrawHeading     bool
        RawJSONData     template.JS
        LanguageStrings map[string]string
    }{
        Title:           lang["map_title"],
        AppName:         AppName,
        AppVersion:      AppVersion,
        GeneratedTime:   htmlFileName,
        DrawHeading:     cfg.ShowCones,
        RawJSONData:     template.JS(jsonData),
        LanguageStrings: lang,
    }

    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, data); err != nil {
        return fmt.Errorf("error executing HTML template: %v", err)
    }

    if err := os.MkdirAll(cfg.OutputDirectory, 0755); err != nil {
        return fmt.Errorf("%s %v", lang["error_creating_output"], err)
    }

    outPath := filepath.Join(cfg.OutputDirectory, htmlFileName)
    if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
        return fmt.Errorf("error writing HTML file: %v", err)
    }

    return nil
}
