package main

type User struct {
	Name  string `yaml:"Name"`
	Email string `yaml:"Email"`
	Alias string `yaml:"alias,omitempty"`
}
