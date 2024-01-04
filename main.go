package main

import (
	"log"
	"os"
)

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Fatalf("reading configuration failed: %s", err)
	}

	cli := NewCLI(cfg)
	if err := cli.App().Run(os.Args); err != nil {
		log.Fatalf("%s", err)
	}

	err = cfg.CloseFile()
	if err != nil {
		log.Fatalf("failed to close the configuration file: %s", err)
	}
}
