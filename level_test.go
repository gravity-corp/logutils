package logutils

import (
	"bytes"
	"log"
	"testing"
)

func TestLevel(t *testing.T) {
	buf := bytes.Buffer{}

	lvl := NewLevel(&buf, "DEBUG")

	log.SetOutput(lvl)
	log.Print("[DEBUG] test")

	if buf.String() == "" {
		t.Error("expected log entry")
	}
}

func TestFilter(t *testing.T) {
	buf := bytes.Buffer{}

	lvl := NewLevel(&buf, "DEBUG")

	log.SetOutput(lvl)
	log.Print("[" + "INFO" + "] test")

	if buf.String() != "" {
		t.Errorf("log entry with the prefix '%s' was not expected", "INFO")
	}
}
