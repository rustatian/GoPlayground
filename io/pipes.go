package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

type augmentedReader struct {
	innerReader io.Reader
	augmentFunc func([]byte) []byte
}

func (r *augmentedReader) Read(buf []byte) (int, error) {
	tmpBuf := make([]byte, len(buf))
	n, err := r.innerReader.Read(tmpBuf)
	copy(buf[:n], r.augmentFunc(tmpBuf[:n]))
	return n, err
}

func encrypt(s []byte) []byte {
	result := make([]byte, len(s))
	for i, c := range s {
		result[i] = c + 28
	}
	return result
}

type augmentedWriter struct {
	innerWriter io.Writer
	augmentFunc func([]byte) []byte
}

func (w *augmentedWriter) Write(buf []byte) (int, error) {
	return w.innerWriter.Write(w.augmentFunc(buf))
}

func EncryptReader(r io.Reader) io.Reader {
	return &augmentedReader{innerReader: r, augmentFunc: encrypt}
}
func UpcaseWriter(w io.Writer) io.Writer {
	return &augmentedWriter{innerWriter: w, augmentFunc: bytes.ToUpper}
}

func main() {
	originalReader := strings.NewReader("this is the stuff I'm reading\n")
	originalWriter := os.Stdout

	pipeReader, pipeWriter := io.Pipe()

	go func() {
		defer pipeWriter.Close()
		_, err := io.Copy(UpcaseWriter(pipeWriter), originalReader)
		if err != nil {
			log.Fatal(err)
		}
	}()

	defer pipeReader.Close()
	_, err := io.Copy(originalWriter, EncryptReader(pipeReader))
	if err != nil {
		log.Fatal(err)
	}
	// output: 'pdeo<eo<pda<opqbb<eCi<na]`ejc&' (notably not uppercased)
}
