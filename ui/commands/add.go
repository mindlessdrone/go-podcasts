package commands

import (
	"errors"
	"fmt"

	"github.com/mindlessdrone/go-podcasts/appl"

	"github.com/urfave/cli"
)

func AddCommand(feedServices *appl.FeedServices) *cli.Command {
	return &cli.Command{
		Name:        "add",
		Description: "adds a podcast feed by url",
		Action: func(c *cli.Context) error {
			url := c.Args().First()

			if url == "" {
				return errors.New("URL is missing")
			}

			err := feedServices.AddFeed(url)
			if err != nil {
				return err
			}

			fmt.Println("OK")
			return nil
		},
	}
}
