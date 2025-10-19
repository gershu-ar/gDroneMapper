package main

import "embed"

//go:embed assets/exiftool.exe assets/exiftool_files/** language.json template.html
var embeddedAssets embed.FS
