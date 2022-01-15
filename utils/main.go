package utils

import (
	"path"
	"path/filepath"
	"strings"
)

func GetFilename(fullpath string) string {
	basename := filepath.Base(fullpath)
	return strings.TrimSuffix(basename, GetExtension(basename))
}

func GetExtension(fullpath string) string {
	return path.Ext(fullpath)
}
