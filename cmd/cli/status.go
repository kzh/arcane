package main

import (
	"context"
	"github.com/kzh/arcane/pkg/arcanepb"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var statusCmd = &cobra.Command{
	Use: "status",
	Short: "health check",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":3333", []grpc.DialOption{grpc.WithInsecure()}...)
		if err != nil {
			zap.L().Error("grpc dial", zap.Error(err))
			return
		}
		defer conn.Close()

		client := arcanepb.NewArcaneClient(conn)
		status, err := client.Status(context.Background(), &arcanepb.StatusRequest{})
		if err != nil {
			zap.L().Error("status", zap.Error(err))
			return
		}

		json, err := protojson.Marshal(status)
		if err != nil {
			zap.L().Error("marshal status", zap.Error(err))
			return
		}

		zap.L().Info("status", zap.String("status", string(json)))
	},
}
