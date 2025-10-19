package main

import (
    "encoding/json"
    "fmt"
)

var lang map[string]string

func loadLanguageFromEmbed(langCode string) error {
    data, err := embeddedAssets.ReadFile("language.json")
    if err != nil {
        return err
    }

    var all map[string]map[string]string
    if err := json.Unmarshal(data, &all); err != nil {
        return err
    }

    selected, ok := all[langCode]
    if !ok {
        return fmt.Errorf("language '%s' not found", langCode)
    }

    lang = selected
    return nil
}
