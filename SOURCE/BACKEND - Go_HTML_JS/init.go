package main

import (
    "io/fs"
    "os"
    "path/filepath"
    "strings"
)

func tempPath(name string) string {
    return filepath.Join(os.TempDir(), "gDroneMapper", name)
}

func deployExifAssetsIfMissing() error {
    baseDir := filepath.Join(os.TempDir(), "gDroneMapper")
    os.MkdirAll(baseDir, 0755)

    exePath := tempPath("exiftool.exe")
    if _, err := os.Stat(exePath); os.IsNotExist(err) {
        data, err := embeddedAssets.ReadFile("assets/exiftool.exe")
        if err != nil {
            return err
        }
        os.WriteFile(exePath, data, 0755)
    }

    exifDir := tempPath("exiftool_files")
    if _, err := os.Stat(exifDir); os.IsNotExist(err) {
        fs.WalkDir(embeddedAssets, "assets/exiftool_files", func(path string, d fs.DirEntry, err error) error {
            rel := strings.TrimPrefix(path, "assets/exiftool_files/")
            target := filepath.Join(exifDir, rel)
            if d.IsDir() {
                return os.MkdirAll(target, 0755)
            }
            data, err := embeddedAssets.ReadFile(path)
            if err != nil {
                return err
            }
            return os.WriteFile(target, data, 0644)
        })
    }

    return nil
}
