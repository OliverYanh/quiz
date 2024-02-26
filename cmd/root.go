package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051" // Address of the gRPC server
)

var rootCmd = &cobra.Command{
	Use:   "quiz-client",
	Short: "A CLI client for interacting with the quiz service", // Short description of the root command
	Run: func(cmd *cobra.Command, args []string) {
		// Show usage if no command provided
		cmd.Usage()
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// dialServer dials the gRPC server and returns a client connection.
func dialServer() (*grpc.ClientConn, error) {
	// Dial the gRPC server with insecure transport credentials
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial server: %v", err)
	}
	return conn, nil
}
