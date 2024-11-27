package main

import (
	"io"
	"log"
	"os"
)

type Printer struct {
	File *os.File
}

func NewPrinter(f *os.File) *Printer {
	return &Printer{File: f}
}

func (p *Printer) FPrint(w io.Writer) (int64, error) {
	n, err := io.Copy(w, p.File)
	log.Printf("%d bytes were printed", n)
	return n, err
}

func provideFile(path string) (*os.File, func(), error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.Printf("close file %v", f.Name())
		if err := f.Close(); err != nil {
			log.Printf("close file fail: %v", err)
		}
	}

	return f, cleanup, nil
}
