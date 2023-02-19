package main

import (
	_ "dev/db"
	. "dev/importer"
	alias "dev/server"
	"fmt"
)

func main() {
	srv := alias.MyServer{}
	srv.Address = "localhost:3000"
	srv.Config = "v1.0"
	Process()
	fmt.Println(srv)
}
