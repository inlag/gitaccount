package main

import (
	"log"
	"os"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatal("reading configuration failed")
	}

	cli := NewCLI(cfg)
	if err := cli.App().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
