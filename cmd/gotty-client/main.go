// +build go1.5

package main

import (
	"fmt"
	"os"
	"path"

	gottyclient "github.com/moul/gotty-client"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
)

var VERSION string

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/gotty-client"
	app.Version = VERSION
	app.Usage = "GoTTY client for your terminal"
	app.ArgsUsage = "GOTTY_URL"
	app.BashComplete = func(c *cli.Context) {
		for _, command := range []string{
			"--debug", "--skip-tls-verify", "--use-proxy-from-env",
			"--v2", "--detach-keys", "--ws-origin", "--help",
			"--generate-bash-completion", "--version",
			"http://user:pass@host:1234/path/\\\\?arg=abcdef\\\\&arg=ghijkl",
			"https://user:pass@host:1234/path/\\\\?arg=abcdef\\\\&arg=ghijkl",
			"http://localhost:8000",
		} {
			fmt.Println(command)
		}
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug, D",
			Usage:  "Enable debug mode",
			EnvVar: "GOTTY_CLIENT_DEBUG",
		},
		cli.BoolFlag{
			Name:   "skip-tls-verify",
			Usage:  "Skip TLS verify",
			EnvVar: "SKIP_TLS_VERIFY",
		},
		cli.BoolFlag{
			Name:   "use-proxy-from-env",
			Usage:  "Use Proxy from environment",
			EnvVar: "USE_PROXY_FROM_ENV",
		},
		cli.StringFlag{
			Name:  "detach-keys",
			Usage: "Key sequence for detaching gotty-client",
			Value: "ctrl-p,ctrl-q",
		},
		cli.BoolFlag{
			Name:   "v2",
			Usage:  "For Gotty 2.0",
			EnvVar: "GOTTY_CLIENT_GOTTY2",
		},
		cli.StringFlag{
			Name:   "ws-origin, w",
			Usage:  "WebSocket Origin URL",
			EnvVar: "GOTTY_CLIENT_WS_ORIGIN",
		},
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "User for Basic Authentication",
			EnvVar: "GOTTY_CLIENT_USER",
		},
	}

	app.Action = action

	if err := app.Run(os.Args); err != nil {
		logrus.Errorf("error: %v", err)
	}
}

func action(c *cli.Context) error {
	if len(c.Args()) != 1 {
		logrus.Fatalf("usage: gotty-client [GoTTY URL]")
	}

	// setting up logrus
	logrus.SetOutput(os.Stderr)
	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// create Client
	url := c.Args()[0]
	client, err := gottyclient.NewClient(url)
	if err != nil {
		logrus.Fatalf("Cannot create client: %v", err)
	}

	if c.Bool("skip-tls-verify") {
		client.SkipTLSVerify = true
	}

	if c.Bool("use-proxy-from-env") {
		client.UseProxyFromEnv = true
	}

	if c.Bool("v2") {
		client.V2 = true
	}

	if wsOrigin := c.String("ws-origin"); wsOrigin != "" {
		client.WSOrigin = wsOrigin
	}

	if detachKey := c.String("detach-keys"); detachKey != "" {
		escapeKeys, err := gottyclient.ToBytes(detachKey)
		if err != nil {
			logrus.Warnf("Invalid escape key: %v", err)
		} else {
			client.EscapeKeys = escapeKeys
		}
	}

	if user := c.String("user"); user != "" {
		client.User = user
		fmt.Print("Password: ")
		password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		if err != nil {
			return cli.NewExitError("Failed to get password from stdin", 2)
		}
		client.Password = string(password)
	}

	// loop
	if err = client.Loop(); err != nil {
		logrus.Fatalf("Communication error: %v", err)
	}
	return nil
}
