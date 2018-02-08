package server

import (
	"context"
	"log"

	proto1 "github.com/carolove/learningforgolang/gosrc/lesson8/types/proto"
	"github.com/cuigh/auxo/net/rpc"
	// "github.com/cuigh/auxo/net/rpc/codec/json"
	"github.com/cuigh/auxo/net/rpc/codec/proto"
	"github.com/cuigh/auxo/net/transport"
)

type HelloService struct{}

func (HelloService) Hello(context.Context, *proto1.HelloRequest) *proto1.HelloResponse {
	ret := &proto1.HelloResponse{}
	ret.Text = "world!"
	return ret
}

func main() {
	s, _ := rpc.Listen(transport.Address{URL: ":9000"})
	// s.Match(json.Matcher, "json")
	//s.Match(jsoniter.Matcher, "json")
	s.Match(proto.Matcher, "proto")
	//s.Use(filter.Test())
	s.RegisterService("Test", TestService{})
	s.RegisterFunc("Test", "Ping", func() string {
		return "pong"
	})
	log.Fatal(s.Serve())
}
