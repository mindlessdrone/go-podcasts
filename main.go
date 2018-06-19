package main

import (
	"fmt"
	"os"

	"github.com/mindlessdrone/go-podcasts/ui/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		*commands.AddCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
