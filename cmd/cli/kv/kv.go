package kv

import (
	"context"
	"github.com/kzh/arcane/pkg/arcanepb"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kv",
		Short: "kv operations",
	}
	cmd.AddCommand(
		getCmd,
		putCmd,
	)

	return cmd
}

var getCmd = &cobra.Command{
	Use:  "get <key>",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":3333", []grpc.DialOption{grpc.WithInsecure()}...)
		if err != nil {
			zap.L().Error("grpc dial", zap.Error(err))
			return
		}
		defer conn.Close()

		client := arcanepb.NewArcaneClient(conn)
		kv, err := client.KVGet(context.Background(), &arcanepb.KVGetRequest{
			Key: args[0],
		})
		if err != nil {
			zap.L().Info("kv get", zap.String("err", err.Error()))
			return
		}

		zap.L().Info("kv", zap.String("key", kv.Key), zap.String("value", kv.Value))
	},
}

var putCmd = &cobra.Command{
	Use:  "put <key> <value>",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial(":3333", []grpc.DialOption{grpc.WithInsecure()}...)
		if err != nil {
			zap.L().Error("grpc dial", zap.Error(err))
			return
		}
		defer conn.Close()

		client := arcanepb.NewArcaneClient(conn)
		_, err = client.KVPut(context.Background(), &arcanepb.KVPutRequest{
			Key:   args[0],
			Value: args[1],
		})
		if err != nil {
			zap.L().Info("kv get", zap.String("err", err.Error()))
			return
		}

		zap.L().Info("kv put", zap.String("state", "success"))
	},
}
