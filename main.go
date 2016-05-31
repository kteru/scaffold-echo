package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kteru/scaffold-echo/api"
	"github.com/kteru/scaffold-echo/db/dummy"
)

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERR] %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func realMain() error {
	listenHost := "127.0.0.1"
	listenPort := 8080
	secret := "secret"

	flagSet := flag.NewFlagSet(filepath.Base(os.Args[0]), flag.ExitOnError)
	flagSet.StringVar(&listenHost, "h", listenHost, "IP Address to listen on")
	flagSet.IntVar(&listenPort, "p", listenPort, "Port number to listen on")
	flagSet.StringVar(&secret, "s", secret, "Secret to encrypt session")
	flagSet.Parse(os.Args[1:])

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	hdl, err := dummy.NewHandler()
	if err != nil {
		return err
	}

	server, err := api.NewServer(&api.ServerConfig{
		ListenHost: listenHost,
		ListenPort: listenPort,
		Logger:     logger,
		Secret:     []byte(secret),
		DBHandler:  hdl,
	})
	if err != nil {
		return err
	}

	if err := server.Serve(); err != nil {
		return err
	}

	return nil
}
