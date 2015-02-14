package bam

// A Flags represents a BAM record's alignment FLAG field.
type Flags uint16

const (
	Paired        Flags = 0x001 // The read is paired in sequencing, no matter whether it is mapped in a pair.
	ProperPair    Flags = 0x002 // The read is mapped in a proper pair.
	Unmapped      Flags = 0x004 // The read itself is unmapped; conflictive with ProperPair.
	MateUnmapped  Flags = 0x008 // The mate is unmapped.
	Reverse       Flags = 0x010 // The read is mapped to the reverse strand.
	MateReverse   Flags = 0x020 // The mate is mapped to the reverse strand.
	Read1         Flags = 0x040 // This is read1.
	Read2         Flags = 0x080 // This is read2.
	Secondary     Flags = 0x100 // Not primary alignment.
	QCFail        Flags = 0x200 // QC failure.
	Duplicate     Flags = 0x400 // Optical or PCR duplicate.
	Supplementary Flags = 0x800 // Supplementary alignment, indicates alignment is part of a chimeric alignment.
)
