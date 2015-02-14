package bam

import (
	"fmt"
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
	var s string
	for i := 0; i < len(c); i++ {
		s += fmt.Sprint(c[i])
	}
	return s
}
