package server

import (
	"fmt"
	"net"
	"os"
)

const (
	msg_length = 1024
)

type Server struct {
	listener net.Listener
}

func NewServer() *Server {
	ret := &Server{}
	listener, err := net.Listen("tcp", ":9080")
	if err != nil {
		os.Exit(-1)
	}
	ret.listener = listener
	return ret
}

func (s *Server) Start() error {
	defer s.listener.Close()

	for {
		conn, err := s.listener.Accept()
		fmt.Println("accept a conn at server")
		if err == nil {
			s.onConn(conn)
		}

	}
}

func (s *Server) onConn(c net.Conn) {
	go func() {
		data := make([]byte, msg_length)
		defer c.Close()

		for {
			n, err := c.Read(data)
			if err != nil {
				fmt.Printf("read message from lotus failed")
				return
			}

			_, err = c.Write(data[0:n])
			if err != nil {
				fmt.Printf("disconnect from lotus server")
				return
			}
		}
	}()
}
