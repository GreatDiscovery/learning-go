package main

// use dial function to establish connection

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}

func TestDialTCP(t *testing.T) {
	addr := "127.0.0.1:80"
	conn, err := net.Dial("tcp", addr)
	checkError(err)

	_, err = conn.Write([]byte("GET /api/v3/get HTTP/1.1\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}

func TestDialUDPServer(t *testing.T) {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}

func TestDialUDPClient(t *testing.T) {
	sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())
}
