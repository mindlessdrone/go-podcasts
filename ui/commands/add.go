package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

func AddCommand() *cli.Command {
	return &cli.Command{
		Name:        "add",
		Description: "adds a podcast feed by url",
		Action: func(c *cli.Context) error {
			fmt.Println("needs to be implemented!")
			return nil
		},
	}
}
