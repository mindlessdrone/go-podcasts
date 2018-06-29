package commands

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/mindlessdrone/go-podcasts/appl"
	"github.com/urfave/cli"
)

func RefreshCommand(feedServices *appl.FeedServices) *cli.Command {
	return &cli.Command{
		Name:        "refresh",
		Description: "refresh podcasts and grab new episodes",
		Action: func(c *cli.Context) error {
			errChan := make(chan error)
			ids := make([]int, 0)
			var wg sync.WaitGroup

			if c.Args().Present() {
				firstArg := c.Args().First()
				podcastID, err := strconv.Atoi(firstArg)
				if err != nil {
					return err
				}

				ids = append(ids, podcastID)
			} else {
				var err error
				ids, err = feedServices.AllFeedIDS()
				if err != nil {
					return err
				}
			}

			// used to monitor when all podcast refreshes are done
			done := make(chan struct{})
			go func() {
				wg.Wait()
				done <- struct{}{}
			}()

			for _, id := range ids {
				go refreshPodcast(feedServices, id, &wg, errChan)
				wg.Add(1)
			}

			for running := true; running; {
				select {
				case err := <-errChan:
					fmt.Fprintln(os.Stderr, err)
				case <-done:
					running = false
				}
			}
			return nil
		},
	}
}

func refreshPodcast(feedServices *appl.FeedServices, id int, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()
	if err := feedServices.RefreshPodcast(id); err != nil {
		errChan <- err
	} else {
		fmt.Println("refreshed podcast: ", id)
	}
}
