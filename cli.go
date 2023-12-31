package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	. "github.com/urfave/cli/v2"
)

type Configer interface {
	GetUsers() []User
	SaveUserInfo(name, email, alias string) error
	RemoveUser(index int) error
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
			c.Add(),
			c.Remove(),
			c.List(),
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
			&IntFlag{Name: "number", Aliases: []string{"n"}},
			&StringFlag{Name: "alias", Aliases: []string{"a"}},
		},
		Action: func(cCtx *Context) error {
			number := cCtx.Int("number")
			alias := cCtx.String("alias")

			users := c.config.GetUsers()

			if alias != "" {
				for _, user := range users {
					if user.Alias == alias {
						return SetGitConfig(user)
					}
				}
			}

			if number != 0 {
				for i, user := range users {
					if i == number {
						return SetGitConfig(user)
					}
				}
			}

			return nil
		},
	}
}

func (c *CLI) Add() *Command {
	return &Command{
		Name:    "add",
		Usage:   "add user",
		Aliases: []string{"a"},
		Flags: []Flag{
			&StringFlag{Name: "name", Aliases: []string{"n"}},
			&StringFlag{Name: "email", Aliases: []string{"e"}},
			&StringFlag{Name: "alias", Aliases: []string{"a"}},
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

			alias := cCtx.String("alias")

			err := c.config.SaveUserInfo(name, email, alias)
			if err != nil {
				return errors.Wrap(err, "failed to save user information")
			}

			return nil
		},
	}
}

func (c *CLI) Remove() *Command {
	return &Command{
		Name:    "remove",
		Usage:   "remove selected user",
		Aliases: []string{"r"},
		Flags: []Flag{
			&IntFlag{Name: "number", Aliases: []string{"n"}},
			&StringFlag{Name: "alias", Aliases: []string{"a"}},
		},
		Action: func(cCtx *Context) error {
			number := cCtx.Int("number")
			alias := cCtx.String("alias")

			users := c.config.GetUsers()

			if alias != "" {
				for i, user := range users {
					if user.Alias == alias {
						err := c.config.RemoveUser(i)
						if err != nil {
							return err
						}
					}
				}
			}

			if number != 0 {
				err := c.config.RemoveUser(number)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
}

func (c *CLI) List() *Command {
	return &Command{
		Name:    "list",
		Usage:   "show list users",
		Aliases: []string{"l"},
		Flags:   []Flag{},
		Action: func(cCtx *Context) error {
			var response strings.Builder

			// todo: add table print
			response.WriteString(fmt.Sprintf("number: Name - Email (alias) \n\n"))
			for i, user := range c.config.GetUsers() {
				line := fmt.Sprintf("%v: %s - %s (%s) \n", i, user.Name, user.Email, user.Alias)
				response.WriteString(line)
			}

			fmt.Println(response.String())
			return nil
		},
	}
}
