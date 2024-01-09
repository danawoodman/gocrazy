package gocrazy

import (
	"os"
	"path/filepath"
	"strings"
)

// ExpandPath expands a path that starts with `~` or `$HOME` to the user's home directory.
func ExpandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, path[2:]), nil
	} else if strings.HasPrefix(path, "$HOME") {
		home := os.Getenv("HOME")
		if home != "" {
			return filepath.Join(home, path[2:]), nil
		}

	}
	return path, nil
}
