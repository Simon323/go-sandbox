package main

import (
	alias "dev/server"
	"fmt"
)

func main() {
	srv := alias.MyServer{}
	srv.Address = "localhost:3000"
	srv.Config = "v1.0"
	fmt.Println(srv)
}
