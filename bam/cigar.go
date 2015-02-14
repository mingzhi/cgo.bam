package bam

import (
	"bytes"
	"fmt"
)

const (
	CigarMatch       byte = 'M' // Alignment match (can be a sequence match or mismatch).
	CigarInsertion   byte = 'I' // Insertion to the reference.
	CigarDeletion    byte = 'D' // Deletion from the reference.
	CigarSkipped     byte = 'N' // Skipped region from the reference.
	CigarSoftClipped byte = 'S' // Soft clipping (clipped sequences present in SEQ).
	CigarHardClipped byte = 'H' // Hard clipping (clipped sequences NOT present in SEQ).
	CigarPadded      byte = 'P' // Padding (silent deletion from padded reference).
	CigarEqual       byte = '=' // Sequence match.
	CigarMismatch    byte = 'X' // Sequence mismatch.
	CigarBack        byte = 'B' // Skip backwards.
)

type CigarOp struct {
	Type byte
	Len  int
}

func (c CigarOp) String() string {
	return fmt.Sprintf("%d%c", c.Len, c.Type)
}

type Cigar []CigarOp

func (c Cigar) String() string {
	if len(c) == 0 {
		return "*"
	}
	var b bytes.Buffer
	for _, co := range c {
		fmt.Fprint(&b, co)
	}
	return b.String()
}
