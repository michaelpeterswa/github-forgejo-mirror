package main

import (
	"context"
	"fmt"
	"os"

	"github.com/michaelpeterswa/github-forgejo-mirror/internal/config"
	"github.com/michaelpeterswa/github-forgejo-mirror/internal/runner"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "github-forgejo-mirror",
		Usage: "Mirror GitHub repositories to Forgejo",
		Flags: config.Flags(),
		Action: func(ctx context.Context, c *cli.Command) error {
			runnr := runner.NewRunner(c)
			return runnr.Run(ctx)
		},
	}

	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
