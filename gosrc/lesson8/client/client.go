package main

import (
	"context"
	"fmt"

	"github.com/cuigh/auxo/net/rpc"
	_ "github.com/cuigh/auxo/net/rpc/codec/json"
	"github.com/cuigh/auxo/net/transport"
)

func main() {
	c, _ := rpc.Dial("json", transport.Address{URL: "127.0.0.1:9000"})

	var s string
	err := c.Call(context.Background(), "Test", "Hello", []interface{}{"auxo"}, &s)
	if err != nil {
		fmt.Println(err)
	}
}
