package bam

/*
#include "bam.h"
*/
import "C"
import (
	"errors"
	"io"

	"unsafe"
)

// Record represents a SAM/BAM record.
type Record struct {
	QName  string // Query template Name.
	Flag   Flags  // bitwise Flag
	RName  string // Reference sequence Name
	Pos    int    // 1-based leftmost mapping Position.
	MapQ   int    // Mapping Quality
	Seq    string // segment Sequence.
	SeqQ   string // ASCII of based Quality plus 33.
	RNext  string // Reference sequence name of the primary alignment of the Next read in the template.
	PNext  int    // Position of the primary alignment of the Next read in the template.
	TLen   int    // signed observed Template Lenth.
	Cigar  Cigar  // CIGAR
	header *Header
	c_bam1 *C.bam1_t
}

func newRecord(h *Header) *Record {
	r := Record{}
	r.c_bam1 = C.bam_init1()
	r.header = h
	return &r
}

func (r *Record) read(b *BGZF) (err error) {
	i := int(C.bam_read1(b.c_BGZF, r.c_bam1))
	switch {
	case i >= 0:
		err = nil
	case i == -1:
		err = io.EOF
	default:
		err = errors.New("bam: something wrong when reading bam file.")
	}

	if err != nil {
		return
	}

	r.parse()

	return nil
}

func (r *Record) parse() {
	// qname
	var c_str *C.char
	c_str = C.get_qname(r.c_bam1)
	r.QName = C.GoString(c_str)
	freeCChars(c_str)
	// flag
	r.Flag = Flags(r.c_bam1.core.flag)
	// rname
	c_str = C.get_rname(r.header.c_bam_hdr, r.c_bam1)
	r.RName = C.GoString(c_str)
	freeCChars(c_str)
	// pos
	r.Pos = int(r.c_bam1.core.pos)
	// mapq
	r.MapQ = int(r.c_bam1.core.qual)
	// seq
	c_str = C.get_seq(r.c_bam1)
	r.Seq = C.GoString(c_str)
	freeCChars(c_str)
	// seqq
	c_str = C.get_seqq(r.c_bam1)
	r.SeqQ = C.GoString(c_str)
	freeCChars(c_str)
	// rnext
	c_str = C.get_rnext(r.header.c_bam_hdr, r.c_bam1)
	r.RNext = C.GoString(c_str)
	freeCChars(c_str)
	// pnext
	r.PNext = int(r.c_bam1.core.mpos)
	// tlen
	r.TLen = int(r.c_bam1.core.isize)

	r.Cigar = r.parseCigar()
}

func freeCChars(c_str *C.char) {
	C.free(unsafe.Pointer(c_str))
}

func (r *Record) Format() string {
	c_str := C.bam_format1(r.header.c_bam_hdr, r.c_bam1)
	str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))
	return str
}

func (r *Record) Destroy() {
	C.bam_destroy1(r.c_bam1)
}

func (r *Record) parseCigar() Cigar {
	var cigar Cigar
	nCigar := int(r.c_bam1.core.n_cigar)
	for i := 0; i < nCigar; i++ {
		var c CigarOp
		c.Len = int(C.cigar_oplen(r.c_bam1, C.int(i)))
		c.Type = byte(C.cigar_opchr(r.c_bam1, C.int(i)))
		cigar = append(cigar, c)
	}
	return cigar
}
