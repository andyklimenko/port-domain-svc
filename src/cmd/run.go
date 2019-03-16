package cmd

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"port-domain-svc/src/config"
	"port-domain-svc/src/proto"
	"port-domain-svc/src/service"
	"port-domain-svc/src/service/storage"
)

var (
	cfg    config.Config
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "run",
		Run: func(cmd *cobra.Command, args []string) {
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

			if serveErr := serv.Serve(ln); serveErr != nil {
				panic(serveErr)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
}
