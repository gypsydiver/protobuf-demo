package main

import (
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/Sirupsen/logrus"
	"github.com/gypsydiver/protobuf-demo/services/heroes/gen"
)

func configureLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

type heroServer struct{}

func (s *heroServer) GetHero(ctx context.Context, hero *gen.Hero) (*gen.Hero, error) {
	log.WithFields(log.Fields{
		"Alias":       hero.GetAlias(),
		"Superpowers": hero.GetSuperpower(),
		"identity":    hero.GetIdentity(),
	}).Info("Got hero request")
	return hero, nil
}

func main() {
	configureLogger()
	addr := "localhost:5000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Listening on: " + addr)

	srv := grpc.NewServer()
	gen.RegisterShieldServer(srv, &heroServer{})
	srv.Serve(lis)
}
