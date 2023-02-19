package main

import (
	client "dev/client"
	http "dev/client/http"
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
	client1 := client.Client1{}
	client1.Config = "v1.1"
	http1 := http.Http1{}
	http1.Name = "Sandbox"
	fmt.Println(srv)
}
