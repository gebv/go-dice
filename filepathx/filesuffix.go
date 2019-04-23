package filepathx

import (
	"path/filepath"
	"strings"
)

// FileNameWithSuffix return file name with suffix (in a file name).
func FileNameWithSuffix(pathFile string, suffix string) string {
	pathFile = strings.TrimSpace(pathFile)
	if len(pathFile) == 0 {
		return ""
	}
	path := filepath.Dir(pathFile)
	ext := filepath.Ext(pathFile)
	spos := len(path) - 1
	if spos > 1 {
		// exists tail of path
		spos += 2
	}
	name := pathFile[spos : len(pathFile)-len(ext)]
	return filepath.Join(path, name+"_"+suffix+ext)
}
