package commands

import (
	"fmt"

	"github.com/mindlessdrone/go-podcasts/appl"
	"github.com/urfave/cli"
)

func ListCommand(feedServices *appl.FeedServices) *cli.Command {
	return &cli.Command{
		Name:        "list",
		Description: "get a list of all subscribed podcasts",
		Action: func(c *cli.Context) error {
			feeds, err := feedServices.AllFeeds()
			if err != nil {
				return err
			}

			for i, feed := range feeds {
				fmt.Printf("%d: %s\t\t%d unplayed\n", i, feed.Title, feed.NumEpisodesUnplayed())
			}
			return nil
		},
	}
}
