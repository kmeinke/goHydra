package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Gateway struct {
		Server  string
		Port    string
		Timeout string
	}
}

var configFile *string = flag.String("c", "config.gcfg", "path to config file")

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	log.Print("Whee! Got Request!")
	io.WriteString(w, "hello, world!\n")
}

func handleOsSignals() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalChannel

		switch sig {
		case os.Interrupt:
			log.Print("Stopping (os.Interrupt)")
		case syscall.SIGTERM:
			log.Print("Stopping (syscall.SIGTERM)")
		default:
			log.Print("Stopping (unknown signal)")
		}
		os.Exit(1)
	}()
}

func loadConfig(c string) Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, c)
	if err != nil {
		log.Fatal("Could't load config file")
	}

	return cfg
}

func main() {
	var err error

	//knt: server is up
	log.Print("Starting")

	//knt:get arguments
	flag.Parse()
	log.Printf("Using config file: %s", *configFile)

	//knt:setup os signal handling to exit the programm
	handleOsSignals()

	//knt: parse config
	cfg := loadConfig(*configFile)

	//knt:set handler
	http.HandleFunc("/hello", HelloServer)

	//knt:listen
	log.Printf("Listening on: %s:%s ...", cfg.Gateway.Server, cfg.Gateway.Port)
	err = http.ListenAndServe(cfg.Gateway.Server+":"+cfg.Gateway.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	os.Exit(0)
}
