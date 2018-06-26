package commands

import (
	"github.com/mindlessdrone/go-podcasts/appl"
	"github.com/urfave/cli"
)

func RefreshCommand(feedServices *appl.FeedServices) *cli.Command {
	return &cli.Command{
		Name:        "refresh",
		Description: "refresh podcasts and grab new episodes",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}
