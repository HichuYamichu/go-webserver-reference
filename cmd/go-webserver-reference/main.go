package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hichuyamichu/go-webserver-reference/app"
	"github.com/hichuyamichu/go-webserver-reference/store"
)

var port = flag.String("port", "3000", "http service port")
var host = flag.String("host", "127.0.0.1", "http service host")

func main() {
	flag.Parse()

	store.ConnectDB()
	srv := app.New()

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	srv.Logger.Fatal(srv.Start(fmt.Sprintf("%s:%s", *host, *port)))
}
