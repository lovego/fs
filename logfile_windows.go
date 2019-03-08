// +build !windows

package fs

import (
	"log"
	"os"
	"path/filepath"
	"sync"
)

type LogFile struct {
	path string
	sync.RWMutex
	*os.File
}

func NewLogFile(path string) (*LogFile, error) {
	l := LogFile{path: path}
	if err := l.open(); err != nil {
		return nil, err
	}
	return &l, nil
}

func (l *LogFile) open() error {
	if err := os.MkdirAll(filepath.Dir(l.path), 0775); err != nil {
		return err
	}
	file, err := os.OpenFile(l.path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	if l.File != nil {
		if err := l.File.Close(); err != nil {
			log.Println(err)
		}
	}
	l.File = file
	return nil
}
