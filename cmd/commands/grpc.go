package commands

import (
	"github.com/guil95/grpc-golang/config/entrypoints/grpc"
	"github.com/spf13/cobra"
	"log"
)

func NewGrpcCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "grpc",
		Short: "Start grpc server",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan error)

			go grpc.RunGrpcServer(quit)

			err := <-quit

			log.Fatal(err)
		},
	}
}
