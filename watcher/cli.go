package watcher

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func RunCLI() {
	var directoryName string

	app := &cli.App{
		Name: "Application to watch changes in current directory",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "watch",
				Aliases:     []string{"w"},
				Usage:       "watch file changes in current directory",
				Destination: &directoryName,
				EnvVars:     []string{"APP_LANG"},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Watching the directory:", directoryName)
			Watch(directoryName)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
