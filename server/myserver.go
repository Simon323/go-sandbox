package server

import "fmt"

type MyServer struct {
	Address string
	Config  string
}

func init() {
	fmt.Println("Server initialized")
}
