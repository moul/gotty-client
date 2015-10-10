package main

import (
	"os"
	"path"

	"github.com/moul/gotty-client"
	"github.com/moul/gotty-client/vendor/github.com/Sirupsen/logrus"
	"github.com/moul/gotty-client/vendor/github.com/codegangsta/cli"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/gotty-client"
	app.Version = VERSION
	app.Usage = "GoTTY client for your terminal"
	app.ArgsUsage = "GOTTY_URL"

	app.Action = Action

	app.Run(os.Args)
}

func Action(c *cli.Context) {
	if len(c.Args()) != 1 {
		logrus.Fatalf("usage: gotty-client [GoTTY URL]")
	}

	url := c.Args()[0]

	// create Client
	client, err := gottyclient.NewClient(url)
	if err != nil {
		logrus.Fatalf("Cannot create client: %v", err)
	}

	// loop
	if err = client.Loop(); err != nil {
		logrus.Fatalf("Communication error: %v", err)
	}
}
