package bam

/*
#include "bam.h"
*/
import "C"

type Header struct {
	c_bam_hdr *C.bam_hdr_t
}

func NewHeader() *Header {
	return &Header{}
}

// Initial bam header.
func (h *Header) Init() {
	h.c_bam_hdr = C.bam_hdr_init()
}

// Read header from a BGZF file.
func (h *Header) Read(b *BGZF) {
	h.c_bam_hdr = C.bam_hdr_read(b.c_BGZF)
}

func (h *Header) Destroy() {
	C.bam_hdr_destroy(h.c_bam_hdr)
}
