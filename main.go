package main

import (
	"fmt"
	"os"

	"github.com/mindlessdrone/go-podcasts/appl"
	"github.com/mindlessdrone/go-podcasts/ui/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	sqlRepository, err := appl.NewSimpleSQLRepository("podcasts.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	feedServices := appl.NewFeedServices(appl.HTTPFeedRetriever{}, sqlRepository)

	app.Commands = []cli.Command{
		*commands.AddCommand(&feedServices),
		*commands.ListCommand(&feedServices),
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
