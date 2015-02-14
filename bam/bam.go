// Package for reading and writinng SAM/BAM format files.
// Using CGO to port htslib.
package bam

/*
#cgo pkg-config:htslib
#include "bam.h"
*/
import "C"

// Bam file.
type File struct {
	name   string // file name.
	b      *BGZF
	header *Header
}

// Open bam file.
// It will also read header.
func OpenFile(name string) *File {
	var f File
	f.name = name
	f.b = NewBGZF(name)
	// read bam header.
	f.header = NewHeader()
	f.header.Read(f.b)
	return &f
}

// Close bam file.
func (f *File) Close() {
	f.b.Close()
}

type BGZF struct {
	c_BGZF *C.BGZF
}

func NewBGZF(file string) *BGZF {
	b := BGZF{}
	b.c_BGZF = C.bgzf_open(C.CString(file), C.CString("r"))
	return &b
}

func (b *BGZF) Close() {
	C.bgzf_close(b.c_BGZF)
}

func (b *File) Read() (*Record, error) {
	r := newRecord(b.header)
	err := r.read(b.b)
	return r, err
}
