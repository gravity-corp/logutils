package logutils

import (
	"bytes"
	"io"
	"unsafe"
)

type Level struct {
	filter map[string]struct{}
	Writer io.Writer
}

// NewLevel returns a new level. Set it as log output
func NewLevel(output io.Writer, prefs ...string) (lvl Level) {
	filter := make(map[string]struct{})

	for _, prefix := range prefs {
		filter[prefix] = struct{}{}
	}

	return Level{filter: filter, Writer: output}
}

func (lvl Level) Write(p []byte) (n int, err error) {
	x := bytes.IndexByte(p, '[')
	if x < 0 {
		return len(p), nil
	}

	y := bytes.IndexByte(p[x:], ']')
	if y < 0 {
		return len(p), nil
	}

	buf := p[x+1 : x+y]
	prefix := *(*string)(unsafe.Pointer(&buf))

	if _, ok := lvl.filter[prefix]; !ok {
		return len(p), nil
	}

	return lvl.Writer.Write(p)
}
