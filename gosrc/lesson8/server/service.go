package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cuigh/auxo/net/rpc"
	"github.com/cuigh/auxo/net/rpc/codec/json"
	"github.com/cuigh/auxo/net/transport"
)

type TestService struct {
}

func (TestService) Hello(ctx context.Context, name string) string {
	fmt.Println(name)
	return "Hello, " + name
}

func main() {
	s, _ := rpc.Listen(transport.Address{URL: ":9000"})
	s.Match(json.Matcher, "json")
	//s.Match(jsoniter.Matcher, "json")
	//s.Match(proto.Matcher, "proto")
	//s.Use(filter.Test())
	s.RegisterService("Test", TestService{})
	s.RegisterFunc("Test", "Ping", func() string {
		return "pong"
	})
	log.Fatal(s.Serve())
}
