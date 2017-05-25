package server

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/gypsydiver/protobuf-demo/services/greeter/gen"
	xContext "golang.org/x/net/context"
	"google.golang.org/grpc"
)

type greeterServer struct{}

func (s *greeterServer) Greet(ctx xContext.Context,
	info *gen.Info) (*gen.Greeting, error) {

	log.WithFields(log.Fields{
		"info": info.String(),
	}).Info("Incoming Request")

	body := generateGreeting(info)

	log.WithField("msg", body).Info("Greeting...")

	return &gen.Greeting{
		Body: body,
	}, nil
}

func generateGreeting(info *gen.Info) (greeting string) {
	greeting = fmt.Sprintf("%s, %s!", info.GetExpression(), info.GetName())
	peeps := info.GetPeeps()
	switch len(peeps) {
	case 0: //do nothing
	case 1:
		greeting += fmt.Sprintf(" Also, a warm %s to %s.", info.GetExpression(),
			peeps[0].GetName())
	default:
		var peepNames string
		for _, v := range peeps[:len(peeps)-1] {
			peepNames += fmt.Sprintf("%s, ", v.GetName())
		}
		//drop last comma
		peepNames = peepNames[:len(peepNames)-1]

		greeting += fmt.Sprintf(" Also, a warm %s to %s and %s.", info.GetExpression(),
			peepNames, peeps[len(peeps)-1].GetName())
	}

	return
}

// Start creates and instance of the greeter service and runs it
func Start() {
	lis, _ := net.Listen("tcp", "localhost:5000")
	srv := grpc.NewServer()
	gen.RegisterGreeterServer(srv, &greeterServer{})
	srv.Serve(lis)
}
