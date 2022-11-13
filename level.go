package logutils

import (
	"bytes"
	"io"
	"unsafe"
)

// Implement the Level and set it as the log output — log.SetOutput()
type Level struct {
	Filter map[string]struct{}
	Writer io.Writer
}

func (lvl *Level) Write(p []byte) (n int, err error) {
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

	if _, ok := lvl.Filter[prefix]; !ok {
		return len(p), nil
	}

	return lvl.Writer.Write(p)
}

// SetFilter is a gallant method to update the level filter
func SetFilter(prefixes ...string) (filter map[string]struct{}) {
	filter = make(map[string]struct{})

	for _, prefix := range prefixes {
		filter[prefix] = struct{}{}
	}

	return filter
}
