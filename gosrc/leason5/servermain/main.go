package main

import (
	"fmt"

	"github.com/carolove/learningforgolang/gosrc/leason5/server"
)

func main() {
	srv := server.NewServer()

	fmt.Printf("Server is ready...\n")
	srv.Start()
}
