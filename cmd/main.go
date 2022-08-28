package main

import (
	"github.com/guil95/grpc-golang/cmd/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "grpc-golang"}

	rootCmd.AddCommand(commands.NewGrpcCommand())
	rootCmd.AddCommand(commands.NewHTTPCommand())
	_ = rootCmd.Execute()
}
