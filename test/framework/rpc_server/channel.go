package rpc_server

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"sync"
)

type messageType uint8

const (
	MessageTypeRequest  = 0x1
	MessageTypeResponse = 0x2
	MessageTypeData     = 0x3
)

var buffers sync.Pool

func (m messageType) String() string {
	switch m {
	case MessageTypeRequest:
		return "request"
	case MessageTypeResponse:
		return "response"
	case MessageTypeData:
		return "data"
	default:
		return "unknown"
	}
}

type messageHeader struct {
	Length   uint32
	StreamID uint32
	Type     messageType
	Flags    uint8
}

type channel struct {
	conn  net.Conn
	bw    *bufio.Writer
	br    *bufio.Reader
	hrbuf [10]byte
	hwbuf [10]byte
}

func newChannel(conn net.Conn) *channel {
	return &channel{
		conn: conn,
		bw:   bufio.NewWriter(conn),
		br:   bufio.NewReader(conn),
	}
}

func (ch *channel) recv() (messageHeader, []byte, error) {
	// 直接初始化一个slice
	mh, err := readMessageHeader(ch.hrbuf[:], ch.br)
	if err != nil {
		return messageHeader{}, nil, nil
	}
	if mh.Length > 4<<20 {
		if _, err := ch.br.Discard(int(mh.Length)); err != nil {
			return mh, nil, fmt.Errorf("failed to discard after receiving oversized message: %w", err)
		}
		return mh, nil, status.Errorf(codes.ResourceExhausted, "message length %v exceed maximum message size of %v", mh.Length, 4<<20)
	}
	var data []byte
	if mh.Length > 0 {
		data = ch.getmbuf(int(mh.Length))
		if _, err := io.ReadFull(ch.br, data); err != nil {
			return messageHeader{}, nil, fmt.Errorf("failed reading message: %w", err)
		}
	}
	return mh, data, nil
}

func (ch *channel) getmbuf(size int) []byte {
	b, ok := buffers.Get().(*[]byte)
	if !ok || cap(*b) < size {
		bb := make([]byte, size)
		b = &bb
	} else {
		*b = (*b)[:size]
	}
	return *b
}

func readMessageHeader(p []byte, r io.Reader) (messageHeader, error) {
	_, err := io.ReadFull(r, p)
	if err != nil {
		return messageHeader{}, err
	}
	return messageHeader{
		// 网络传输一般使用大端序，计算机内部处理一般用小端序
		Length:   binary.BigEndian.Uint32(p[:4]),
		StreamID: binary.BigEndian.Uint32(p[4:8]),
		Type:     messageType(p[8]),
		Flags:    p[9],
	}, nil
}
