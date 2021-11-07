package main

import (
	"fmt"
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
	c.AddCommand(statusCmd)

	if err := c.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
