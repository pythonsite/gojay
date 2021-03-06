package gojay

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderBorrowFromPool(t *testing.T) {
	// reset pool
	decPool = make(chan *Decoder, 16)
	// borrow decoder
	dec := BorrowDecoder(strings.NewReader(""))
	// release
	dec.Release()
	// get from pool
	nDec := BorrowDecoder(strings.NewReader(""))
	// assert same
	assert.Equal(t, dec, nDec, "both decoders should be the same")
}

func TestDecoderBorrowFromPoolSetBuffSize(t *testing.T) {
	dec := borrowDecoder(nil, 512)
	assert.Len(t, dec.data, 512, "data buffer should be of len 512")
}
