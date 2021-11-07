package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kzh/arcane/pkg/arcanepb"
	"github.com/kzh/arcane/server"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use: "arcane",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":3333")
		if err != nil {
			zap.L().Error("listen", zap.Error(err))
			return
		}

		s := server.NewServer()
		grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
		arcanepb.RegisterArcaneServer(grpcServer, s)

		zap.L().Info("Starting Arcane...", zap.String("foo", "bar"))

		if err := grpcServer.Serve(lis); err != nil {
			zap.L().Error("grpc serve", zap.Error(err))
		}
	},
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
