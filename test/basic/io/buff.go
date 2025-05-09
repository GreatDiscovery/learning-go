package io

import (
	"github.com/pkg/errors"
	"io"
)

const (
	BuffSizeAlign = 1024 * 4
)

type memBuffer struct {
	b    []byte
	size uint64
	rpos uint64
	wpos uint64
}

func newMemBuffer(buffSize int) *memBuffer {
	n := align(buffSize, BuffSizeAlign)
	if n <= 0 {
		panic("invalid pipe buffer size")
	}
	return &memBuffer{b: make([]byte, n), size: uint64(n)}
}

func (p *memBuffer) readSome(b []byte) (int, error) {
	if p.b == nil {
		return 0, errors.WithStack(io.ErrClosedPipe)
	}
	maxlen, offset := roffset(len(b), p.size, p.rpos, p.wpos)
	if maxlen == 0 {
		return 0, nil
	}
	n := copy(b, p.b[offset:offset+maxlen])
	p.rpos += uint64(n)
	if p.rpos == p.wpos {
		p.rpos = 0
		p.wpos = 0
	}
	return n, nil
}

func (p *memBuffer) writeSome(b []byte) (int, error) {
	if p.b == nil {
		return 0, errors.WithStack(io.ErrClosedPipe)
	}
	maxlen, offset := woffset(len(b), p.size, p.rpos, p.wpos)
	if maxlen == 0 {
		return 0, nil
	}
	n := copy(p.b[offset:offset+maxlen], b)
	p.wpos += uint64(n)
	return n, nil
}

func (p *memBuffer) buffered() int {
	if p.b == nil {
		return 0
	}
	return int(p.wpos - p.rpos)
}

func (p *memBuffer) available() int {
	if p.b == nil {
		return 0
	}
	return int(p.size + p.rpos - p.wpos)
}

func (p *memBuffer) rclose() error {
	p.b = nil
	return nil
}

func (p *memBuffer) wclose() error {
	return nil
}
