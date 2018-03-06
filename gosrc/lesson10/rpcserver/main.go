package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cuigh/auxo/net/rpc"
	"github.com/cuigh/auxo/net/transport"
)

type TestService struct {
}

func (TestService) Hello(ctx context.Context, name string) string {
	return "Hello, " + name
}

func main() {
	s, err := rpc.Listen(transport.Address{URL: ":9000"})
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Match(rpc.Any, "json")
	s.RegisterService("Test", TestService{})
	s.RegisterFunc("Test", "Ping", func() string {
		return "pong"
	})
	log.Fatal(s.Serve())
}
