package bam

/*
#cgo pkg-config:htslib
#include "bam.h"
*/
import "C"
import (
	"errors"
	"io"
)

// Bam file.
type File struct {
	name   string // file name.
	b      *BGZF
	header *Header
}

// Open bam file.
func OpenFile(name string) *File {
	var f File
	f.name = name
	f.b = NewBGZF(name)
	// read bam header.
	f.header = NewHeader()
	f.header.Read(f.b)
	n := int(f.header.c_bam_hdr.n_targets)
	println(n)
	println(C.GoString(f.header.c_bam_hdr.text))
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
	var err error
	r := newRecord()
	r.header = b.header
	i := int(C.bam_read1(b.b.c_BGZF, r.c_bam1))
	switch {
	case i >= 0:
		err = nil
	case i == -1:
		err = io.EOF
	default:
		err = errors.New("bam: something wrong when reading bam file.")
	}
	return r, err
}
