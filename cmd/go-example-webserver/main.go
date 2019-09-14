package main

import (
	"flag"

	"github.com/hichuyamichu/go-webserver-reference/app"
)

var port = flag.String("port", "3000", "http service port")
var host = flag.String("host", "127.0.0.1", "http service host")

func main() {
	flag.Parse()
	srv := app.New(*host, *port)
	srv.Run()
}
