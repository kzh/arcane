package main

import (
	"fmt"
	"github.com/kzh/arcane/cmd/cli/kv"
	"github.com/kzh/arcane/cmd/cli/status"
	"go.uber.org/zap"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
		return
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	c := &cobra.Command{
		Use: "arcanecli",
	}
	c.AddCommand(
		status.Cmd(),
		kv.Cmd(),
	)

	if err := c.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
