package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gogoscan/engine"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "goscan",
		Usage: "my study gogoscan service",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "ip",
				Usage:   "attack target ip",
				Required:true,
			},
			&cli.StringFlag{
				Name:    "protocol",
				Aliases: []string{"p"},
				Usage:   "protocol support ssh",
				Value: "ssh",
				Required: true,
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
				Value:       20,
			},
			&cli.StringFlag{
				Name:        "username_path",
				Aliases:     []string{"username"},
				Usage:       "username text path",
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "password_path",
				Aliases:     []string{"password"},
				Usage:       "password text path",
				Required:    false,
			},
		},
		Action: func(c *cli.Context) error {
			scanEngine := engine.ScanEngine{
					Ip:               c.String("ip"),
					Port:             c.Int("port"),
					Protocol:         c.String("protocol"),
					UsernameTextPath: c.String("username_path"),
					PasswordTextPath: c.String("password_path"),
					Concurrency:      c.Int("concurrency"),
			}

			if scanEngine.UsernameTextPath == ""{
				scanEngine.UsernameTextPath = "./resource/username.txt"
			}
			if scanEngine.PasswordTextPath == "" {
				scanEngine.PasswordTextPath = "./resource/password.txt"
			}
			scanEngine.Run()
			fmt.Println("Result: ", scanEngine.PasswordBurst)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}