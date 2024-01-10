package gocrazy_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/danawoodman/gocrazy"
)

func TestExpandHomeTilde(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	val, err := gocrazy.ExpandHome("~/test")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if val != filepath.Join(home, "test") {
		t.Errorf("Expected %s, got %s", filepath.Join(home, "test"), val)
	}
}

func TestExpandHomeDollarHome(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	val, err := gocrazy.ExpandHome("$HOME/test")
	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if val != filepath.Join(home, "test") {
		t.Errorf("Expected %s, got %s", filepath.Join(home, "test"), val)
	}
}
