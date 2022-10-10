package main

import "preco/client"

func main() {
	client := &client.Client {Port1: "8000", Port2: "8080", Addr1: "localhost", Addr2: "localhost"}
	
	client.Write(2, []uint64{1, 2, 3, 4, 5})
}