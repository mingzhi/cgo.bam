package bam

import (
	"bytes"
	"fmt"
)

const (
	CigarMatch       byte = 'M'
	CigarInsertion   byte = 'I'
	CigarDeletion    byte = 'D'
	CigarSkipped     byte = 'N'
	CigarSoftClipped byte = 'S'
	CigarHardClipped byte = 'H'
	CigarPadded      byte = 'P'
	CigarEqual       byte = '='
	CigarMismatch    byte = 'X'
	CigarBack        byte = 'B'
)

type CigarOp struct {
	C   byte
	Len int
}

func (c CigarOp) String() string {
	return fmt.Sprintf("%d%c", c.Len, c.C)
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
