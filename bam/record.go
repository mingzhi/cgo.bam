package bam

/*
#include "sam.h"
*/
import "C"

type Record struct {
	c_bam1_t *C.bam1_t
}

func newRecord() *Record {
	r := Record{}
	r.c_bam1_t = C.bam_init1()
	return &r
}

func (r *Record) Pos() int {
	return int(r.c_bam1_t.core.pos)
}
