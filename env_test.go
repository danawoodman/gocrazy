package gocrazy_test

import (
	"testing"

	"github.com/danawoodman/gocrazy"
)

func TestGetEnvNonExistent(t *testing.T) {
	val := gocrazy.Getenv("THIS_DOESNT_EXIST", "default")
	if val != "default" {
		t.Errorf("Expected default value, got %s", val)
	}
}

func TestGetEnvExists(t *testing.T) {
	fallback := "should_not_get_here"
	val := gocrazy.Getenv("GO111MODULE", fallback)
	if val == fallback {
		t.Errorf("Expected GO111MODULE env var to be set, got %s", val)
	}
}
