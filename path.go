package gocrazy

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandHome expands a path that starts with `~` or `$HOME` to the user's home directory.
func ExpandHome(path string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(path, "~/") {
		return filepath.Join(home, path[2:]), nil
	} else if strings.HasPrefix(path, "$HOME") {
		if home != "" {
			return filepath.Join(home, path[5:]), nil
		}

	}
	return path, nil
}
