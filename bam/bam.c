#include "bam.h"

char * bam_format1(const bam_hdr_t *header, const bam1_t *b)
{
    kstring_t str;
    str.l = str.m = 0; str.s = NULL;
    sam_format1(header, b, &str);
    return ks_release(&str);
}

char * get_qname(const bam1_t *b) {
	const bam1_core_t *c = &b->core;
	kstring_t str;
	str.l = str.m = 0; str.s = NULL;
	kputsn(bam_get_qname(b), c->l_qname-1, &str);
	return ks_release(&str);
}

char * get_rname(const bam_hdr_t *h, const bam1_t *b) {
	const bam1_core_t *c = &b->core;
	kstring_t str;
	str.l = str.m = 0; str.s = NULL;
	if (c->tid >= 0)
	{
		kputs(h->target_name[c->tid], &str);
	} else {
		kputsn("*", 1, &str);
	}
	return ks_release(&str);
}

char * get_seq(const bam1_t *b) {
	const bam1_core_t *c = &b->core;
	int i;
	kstring_t str;
	str.l = str.m = 0; str.s = NULL;
	if (c->l_qseq) { // seq and qual
        uint8_t *s = bam_get_seq(b);
        for (i = 0; i < c->l_qseq; ++i) kputc("=ACMGRSVTWYHKDBN"[bam_seqi(s, i)], &str);
    } else kputsn("*", 1, &str);
	return ks_release(&str);
}

char * get_seqq(const bam1_t *b) {
	const bam1_core_t *c = &b->core;
	int i;
	kstring_t str;
	str.l = str.m = 0; str.s = NULL;
	if (c->l_qseq) { // seq and qual
        uint8_t *s = bam_get_qual(b);
        if (s[0] == 0xff) kputc('*', &str);
        else for (i = 0; i < c->l_qseq; ++i) kputc(s[i] + 33, &str);
    } else kputsn("*", 1, &str);
	return ks_release(&str);
}