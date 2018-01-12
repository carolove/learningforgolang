package server

import (
	"encoding/json"
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
		if err == nil {
			go s.onConn(conn)
		}

	}
}

func (s *Server) onConn(c net.Conn) {
	data := make([]byte, msg_length)
	defer c.Close()

	var recvdata map[string]string
	recvdata = make(map[string]string, 2)
	var senddata map[string]string
	senddata = make(map[string]string, 2)
	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Printf("read message from lotus failed")
			return
		}

		if err := json.Unmarshal(data[0:n], &recvdata); err == nil {
			senddata["reqId"] = recvdata["reqId"]
			senddata["resContent"] = "Hello " + recvdata["reqContent"]

			sendjson, err := json.Marshal(senddata)
			_, err = c.Write([]byte(sendjson))
			if err != nil {
				fmt.Printf("disconnect from lotus server")
				return
			}
		}
	}
}
