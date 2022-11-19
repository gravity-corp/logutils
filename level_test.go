package logutils

import (
	"bytes"
	"log"
	"testing"
)

func TestLevel(t *testing.T) {
	buf := bytes.Buffer{}

	lvl := Level{
		Filter: map[string]struct{}{"DEBUG": {}},
		Writer: &buf,
	}

	log.SetOutput(lvl)
	log.Print("[DEBUG] test")

	if buf.String() == "" {
		t.Error("Expected log entry")
	}
}

func TestFilter(t *testing.T) {
	prefix1, prefix2 := "DEBUG", "INFO"
	buf := bytes.Buffer{}

	lvl := &Level{
		Filter: map[string]struct{}{prefix1: {}},
		Writer: &buf,
	}

	log.SetOutput(lvl)
	log.Print("[" + prefix2 + "] test")

	if buf.String() != "" {
		t.Errorf("Log entry with the prefix '%s' was not expected", prefix2)
	}
}

func TestSetFilter(t *testing.T) {
	prefix := "DEBUG"
	filter := SetFilter(prefix)

	if _, ok := filter[prefix]; !ok {
		t.Errorf("The prefix '%s' was expected to be set", prefix)
	}
}
