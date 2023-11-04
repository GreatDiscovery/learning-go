package rpc_server

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"testing"
)

func TestPb(t *testing.T) {
	request := Request{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		Service:       "service1",
		Method:        "method1",
		Payload:       nil, // Payload里可以放其他的pb解析后的byte[]
		TimeoutNano:   0,
		Metadata:      nil,
	}
	marshal, err := proto.Marshal(&request)
	if err != nil {
		fmt.Println(err)
		return
	}
	var reqeust2 Request
	err = proto.Unmarshal(marshal, &reqeust2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request.String())
	fmt.Println(reqeust2.String())
}
