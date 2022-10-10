package main

import "preco/server"

func main() {
	db := server.NewDatabase([][]uint64{{1,2,3}, {4, 5, 6}})
	s := &server.Server{Database: db}
	s.RunServer("8080")
}