package main

import (
	"flag"
	"fmt"

	"github.com/nepp-tumsat/documents-api/server"
)

var addr string

func init() {
	port := "8000"

	flag.StringVar(&addr, "addr", fmt.Sprintf(":%s", port), "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}
