package rpc_server

import (
	"bufio"
	"net"
)

type channel struct {
	conn net.Conn
	bw   *bufio.Writer
	br   *bufio.Reader
}

func newChannel(conn net.Conn) *channel {
	return &channel{
		conn: conn,
		bw:   bufio.NewWriter(conn),
		br:   bufio.NewReader(conn),
	}
}
