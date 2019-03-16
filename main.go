package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"ports/port-domain-svc/src/config"
	"ports/port-domain-svc/src/proto"
	"ports/port-domain-svc/src/service"
	"ports/port-domain-svc/src/service/storage"
)

func main() {
	var cfg config.Config
	cfg.Load()

	if cfg.Port == "" {
		panic("grpc server port not specified")
	}

	ln, lnErr := net.Listen("tcp", cfg.Port)
	if lnErr != nil {
		panic(lnErr)
	}

	st, stErr := storage.NewStorage(cfg.Pg)
	if stErr != nil {
		panic(stErr)
	}
	serv := grpc.NewServer()
	portdomainsvc.RegisterPortDomainServiceServer(serv, service.NewService(st))
	reflection.Register(serv)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)
	go func() {
		<-sigCh
		serv.Stop()
	}()

	if err := serv.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
