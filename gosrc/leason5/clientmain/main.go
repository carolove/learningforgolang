package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/carolove/learningforgolang/gosrc/leason5/client"
)

func main() {
	var tc client.Clienter
	tc.SendStr = make(chan string, 1)
	tc.RecvStr = make(chan string, 1)
	if !tc.Connect() {
		return
	}
	r := bufio.NewReader(os.Stdin)
	for {
		switch line, ok := r.ReadString('\n'); true {
		case ok != nil:
			fmt.Printf("bye bye!\n")
			return
		default:
			go client.Work(&tc)
			tc.SendStr <- line
			s := <-tc.RecvStr
			fmt.Printf("back:%s\n", s)
		}
	}
}
