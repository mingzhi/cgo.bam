package bam

/*
#cgo pkg-config:htslib
#include "sam.h"
*/
import "C"
import (
	"errors"
	"io"
)

type BGZF struct {
	c_BGZF *C.BGZF
}

type BamFile struct {
	file string
	b    *BGZF
}

func NewBGZF(file string) *BGZF {
	b := BGZF{}
	b.c_BGZF = C.bgzf_open(C.CString(file), C.CString("r"))
	return &b
}

func (b *BGZF) Close() {
	C.bgzf_close(b.c_BGZF)
}

func NewBamFile(file string) *BamFile {
	b := BamFile{}
	b.file = file
	b.b = NewBGZF(file)
	return &b
}

func (b *BamFile) Close() {
	b.b.Close()
}

func (b *BamFile) Read() (*Record, error) {
	var err error
	r := newRecord()
	i := C.bam_read1(b.b.c_BGZF, r.c_bam1_t)
	switch i {
	case 0:
		err = nil
	case -1:
		err = io.EOF
	default:
		err = errors.New("bam: something wrong when reading bam file.")
	}
	return r, err
}
