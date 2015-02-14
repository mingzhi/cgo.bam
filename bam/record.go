package bam

/*
#include "bam.h"
*/
import "C"
import "unsafe"

type Record struct {
	header *Header
	c_bam1 *C.bam1_t
}

func newRecord() *Record {
	r := Record{}
	r.c_bam1 = C.bam_init1()
	return &r
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

func (r *Record) QName() string {
	c_str := C.get_qname(r.c_bam1)
	str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))
	return str
}

func (r *Record) Flag() int {
	return (int(r.c_bam1.core.flag))
}

func (r *Record) RName() string {
	c_str := C.get_rname(r.header.c_bam_hdr, r.c_bam1)
	str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))
	return str
}

func (r *Record) Pos() int {
	return (int(r.c_bam1.core.pos) + 1)
}

func (r *Record) MapQ() int {
	return (int(r.c_bam1.core.qual))
}

func (r *Record) Seq() string {
	c_str := C.get_seq(r.c_bam1)
	str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))
	return str
}

func (r *Record) SeqQ() string {
	c_str := C.get_seqq(r.c_bam1)
	str := C.GoString(c_str)
	C.free(unsafe.Pointer(c_str))
	return str
}
