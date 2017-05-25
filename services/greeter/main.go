package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gypsydiver/protobuf-demo/services/greeter/server"
)

func configureLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	configureLogger()
	server.Start()
}
