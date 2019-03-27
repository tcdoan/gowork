//!+

// Package bzip provides a writter
// that uses bzip2 compression bzip.org
//
package bzip

/*
#cgo CFLAGS: -I/usr/include
#cgo LDFLAGS: -L/usr/lib -lbz2
#include  <bzlib.h>
#include  <stdlib.h>
bz_stream* bz2alloc() { return calloc(1, sizeof(bz_stream)); }
int bz2compress(
    bz_stream *s,
    int action,
    char * in,
    unsigned *inlen,
    char* out,
    unsigned *outlen);
void bz2free(bz_stream *s) { free(s); }
*/
import "C"

import (
	"io"
	"unsafe"
)

type writter struct {
	w      io.Writer
	stream *C.bz_stream
	outbuf [64 * 1024]byte
}

// Bz2Writter returns a writter for bzip2-compressed streams
func Bz2Writter(out io.Writer) io.WriteCloser {
	const (
		blockSize  = 9
		verbosity  = 0
		workFactor = 30
	)
	w := &writter{w: out, stream: C.bz2alloc()}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

//!+write
func (w *writter) Write(data []byte) (int, error) {
	if w.stream == nil {
		panic("close")
	}

	var total int // uncompressed bytes written

	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(
			w.stream,
			C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])),
			&inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)),
			&outlen)
		total += int(inlen)
		data = data[inlen:]

		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}
	return total, nil
}

//!-close
func (w *writter) Close() error {
	if w.stream == nil {
		panic("close")
	}

	defer func() {
		C.BZ2_bzCompressEnd(w.stream)
		C.bz2free(w.stream)
		w.stream = nil
	}()

	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(
			w.stream,
			C.BZ_FINISH,
			nil,
			&inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)),
			&outlen,
		)

		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}

		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}
