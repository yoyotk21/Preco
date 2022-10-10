package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Server struct {
	Database *Database
}

func (server *Server) RunServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port", port)

	// Closing the connection once the program is finished
	defer listener.Close()

	rpcServer := rpc.NewServer()

	rpcServer.Register(server.Database)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpcServer.ServeConn(conn)
	}
}
