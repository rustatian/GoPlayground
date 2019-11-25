package main

import (
	"bytes"
)

//go:noinline
func PackWoOptimization(m string, s uint64) []byte {
	b := bytes.Buffer{}
	b.WriteString(m)

	b.Write([]byte{
		byte(s),
		byte(s >> 8),
		byte(s >> 16),
		byte(s >> 24),
		byte(s >> 32),
		byte(s >> 40),
		byte(s >> 48),
		byte(s >> 56),
	})

	return b.Bytes()
}

//go:noinline
func PackWOptimization(m string, s uint64, buf []byte) {
	copy(buf[0:], m)
	copy(buf[len(m):], []byte{
		byte(s),
		byte(s >> 8),
		byte(s >> 16),
		byte(s >> 24),
		byte(s >> 32),
		byte(s >> 40),
		byte(s >> 48),
		byte(s >> 56),
	})
}
