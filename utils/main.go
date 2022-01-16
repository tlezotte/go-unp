package utils

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/gookit/color"
)

var (
	Fade   = color.FgGray.Render
	Red    = color.FgRed.Render
	Danger = color.Danger.Render
	Info   = color.Info.Render
	Error  = color.New(color.FgRed, color.OpBold)
	Green  = color.FgGreen.Render
)

func GetFilename(fullpath string) string {
	basename := filepath.Base(fullpath)
	return strings.TrimSuffix(basename, GetExtension(basename))
}

func GetExtension(fullpath string) string {
	return path.Ext(fullpath)
}

func Contains(extList []string, searchterm string) bool {
	for _, a := range extList {
		if a == searchterm {
			return true
		}
	}
	return false
}
