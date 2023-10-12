package rpc_server

import (
	"bufio"
	"encoding/binary"
	"io"
	"net"
)

type messageType uint8

const (
	MessageTypeRequest  = 0x1
	MessageTypeResponse = 0x2
	MessageTypeData     = 0x3
)

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
	StreamId uint32
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
	return mh, nil, nil
}

func readMessageHeader(p []byte, r io.Reader) (messageHeader, error) {
	_, err := io.ReadFull(r, p)
	if err != nil {
		return messageHeader{}, err
	}
	return messageHeader{
		// 网络传输一般使用大端序，计算机内部处理一般用小端序
		Length:   binary.BigEndian.Uint32(p[:4]),
		StreamId: binary.BigEndian.Uint32(p[4:8]),
		Type:     messageType(p[8]),
		Flags:    p[9],
	}, nil
}
