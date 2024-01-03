package main

import (
	"fmt"

	"github.com/pkg/errors"
	. "github.com/urfave/cli/v2"
)

type Configer interface {
	GetUsers() []User
	SaveUserInfo(name string, email string) error
}

type CLI struct {
	config Configer
}

func NewCLI(cfg *Config) *CLI {
	return &CLI{
		config: cfg,
	}
}

func (c *CLI) App() *App {
	return &App{
		Name:                 "gitacc",
		EnableBashCompletion: true,
		Commands: []*Command{
			c.Select(),
		},
		Usage: "managing multiple git accounts",
		Action: func(*Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}
}

func (c *CLI) Select() *Command {
	return &Command{
		Name:    "select",
		Usage:   "user selection",
		Aliases: []string{"s"},
		Flags: []Flag{
			&StringFlag{Name: "name", Aliases: []string{"n"}},
			&StringFlag{Name: "email", Aliases: []string{"e"}},
		},
		Action: func(cCtx *Context) error {
			name := cCtx.String("name")
			if name == "" {
				return errors.New("name cannot is empty")
			}

			email := cCtx.String("email")
			if email == "" {
				return errors.New("email cannot is empty")
			}

			err := c.config.SaveUserInfo(name, email)
			if err != nil {
				return errors.Wrap(err, "failed to save user information")
			}

			return nil
		},
	}
}
