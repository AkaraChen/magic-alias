package main

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

func init() {}

func main() {
	app := cli.Command{
		Name:        "magic-alias",
		Description: "Alias, but magical.",
	}

	app.Run(context.Background(), os.Args)
}
