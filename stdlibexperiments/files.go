package stdlibexperiments

import (
	"os"
	"sync"
	"time"
)

func CreateFile(name string) *FileData {
	os.Create()

	return &FileData{name: name, mode: os.ModeTemporary, modtime: time.Now()}
}

type File struct {
	at           int64
	readDirCount int64
	closed       bool
	readOnly     bool
	fileData     *FileData
}

type FileData struct {
	sync.Mutex
	name    string
	data    []byte
	memDir  Dir
	dir     bool
	mode    os.FileMode
	modtime time.Time
}

type Dir interface {
	Len() int
	Names() []string
	Files() []*FileData
	Add(*FileData)
	Remove(*FileData)
}
