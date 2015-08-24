package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/moul/gotty-client"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [GoTTY URL]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		logrus.Fatalf("GoTTY URL is missing.")
	}

	// create Client
	client, err := gottyclient.NewClient(flag.Arg(0))
	if err != nil {
		logrus.Fatalf("Cannot create client: %v", err)
	}

	// loop
	err = client.Loop()
	if err != nil {
		logrus.Fatalf("Communication error: %v", err)
	}
}
