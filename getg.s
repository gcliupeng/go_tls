#ifdef GOARCH_amd64
#define	get_tls(r)	MOVQ TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

#ifdef GOARCH_amd64p32
#define	get_tls(r)	MOVL TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

#ifdef GOARCH_386
#define	get_tls(r)	MOVL TLS, r
#define	g(r)	0(r)(TLS*1)
#endif

TEXT go_tls·getg(SB),$0-8
    get_tls(BX)
    MOVQ g(BX), CX
    MOVQ CX, ret+0(FP)
    RET

