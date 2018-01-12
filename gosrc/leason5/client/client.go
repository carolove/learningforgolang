package client

import (
	"fmt"
	"net"
	"time"
)

type Clienter struct {
	client  net.Conn
	isAlive bool
	SendStr chan string
	RecvStr chan string
}

func (c *Clienter) Connect() bool {
	if c.isAlive {
		return true
	} else {
		var err error
		c.client, err = net.Dial("tcp", "127.0.0.1:9080")
		if err != nil {
			fmt.Printf("Failure to connet:%s\n", err.Error())
			return false
		}
		c.isAlive = true
	}
	return true
}

func (c *Clienter) Echo() {
	line := <-c.SendStr
	c.client.Write([]byte(line))
	buf := make([]byte, 1024)
	n, err := c.client.Read(buf)
	if err != nil {
		c.RecvStr <- string("Server close...")
		c.client.Close()
		c.isAlive = false
		return
	}
	time.Sleep(1 * time.Second)
	c.RecvStr <- string(buf[0:n])
}

func Work(tc *Clienter) {
	if !tc.isAlive {
		if tc.Connect() {
			tc.Echo()
		} else {
			<-tc.SendStr
			tc.RecvStr <- string("Server close...")
		}
	} else {
		tc.Echo()
	}
}
