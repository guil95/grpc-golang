package commands

import (
	"github.com/guil95/grpc-golang/config/entrypoints/http"
	"github.com/spf13/cobra"
	"log"
)

func NewHTTPCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "http",
		Short: "Start http server",
		Run: func(cmd *cobra.Command, args []string) {
			quit := make(chan error)

			go http.RunHTTPServer(quit)

			err := <-quit

			log.Fatal(err)
		},
	}
}
