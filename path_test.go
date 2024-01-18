package gocrazy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danawoodman/gocrazy"
)

func TestExpandHome(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("Expected to finder user homed directory: %s", err)
	}

	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"can handle tilde (~/) home dir", "~/test", filepath.Join(home, "test")},
		{"can handle dollar home ($HOME) home dir", "$HOME/test", filepath.Join(home, "test")},
		{"does nothing on non-home root path", "/foo/bar", "/foo/bar"},
		{"does nothing on non-home relative path", "./foo/bar", "./foo/bar"},
		{"does nothing on blank string", "", ""},
	}

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			val, err := gocrazy.ExpandHome(tt.input)
			if err != nil {
				t.Errorf("Expected not to error, got error: %s", err)
			}

			if val != tt.want {
				t.Errorf("Expected %#v, got %#v", tt.want, tt.input)
			}
		})
	}
}
