package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "goscan",
		Usage: "my study gogoscan service",
		Action: func(c *cli.Context) error {
			fmt.Println(`
Support:
  [x] ssh-bruturs:	support ssh brutus attack 
`)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "ssh-brutus",
				Aliases: []string{"ssh-brutus"},
				Usage:   "ssh brutus force attack",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "ip",
						Usage:   "attack target ip",
						Required:true,
					},
					&cli.IntFlag{
						Name:        "port",
						Usage: "attack target port",
						Required:    true,
					},
					&cli.IntFlag{
						Name:        "concurrency",
						Aliases:     []string{"c"},
						Usage:       "the threads of concurrency num",
						Required:    false,
						Value:       50,
					},
				},
				Action: func(c *cli.Context) error {
					SSHScanStart(c.String("ip"), c.Int("port"), c.Int("concurrency"))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}