#ifndef SAM_H
#define SAM_H
#include <htslib/bgzf.h>
#include <htslib/sam.h>
#include <htslib/kstring.h>

char *bam_format1(const bam_hdr_t *header, const bam1_t *b);
char * get_qname(const bam1_t *b);
char * get_rname(const bam_hdr_t *h, const bam1_t *b);
char * get_seq(const bam1_t *b);
char * get_seqq(const bam1_t *b);
int cigar_oplen(const bam1_t *b, int i);
char cigar_opchr(const bam1_t *b, int i);
char * get_rnext(const bam_hdr_t *h, const bam1_t *b);
#endif