package main

import (
	"fmt"
	"log"
	"os"

	"devm/indexer/indexer/api"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "indexer",
		HelpName: "indexer",
		Usage:    "Chain Indexer HTTP Server",
		Version:  api.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "`path` to config file",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("action")
			app, err := api.NewIndexer(c.String("config"))
			if err != nil {
				return err
			}
			fmt.Println("preshutdown")
			go app.Shutdown()
			fmt.Println("prestart")

			return app.Start()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		return
	}
}
