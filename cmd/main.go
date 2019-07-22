package main

import (
	"flag"

	"github.com/HichuYamichu/go-webserver-reference/app"
)

var port = flag.String("port", "3000", "http service port")
var mongoURI = flag.String("mongoURI", "mongodb://localhost:27017", "mongoDB instance URI")

func main() {
	srv := app.NewServer(*port, *mongoURI)
	srv.Run()
}
