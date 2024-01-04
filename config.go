package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	f     *os.File
	users []User `yaml:"users"`
}

func NewConfig() (*Config, error) {
	file, err := GetConfigFile()
	if err != nil {
		return nil, errors.Wrap(err, "failed is open configuration file")
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "reading configuration file is failed")
	}

	var users []User
	err = yaml.Unmarshal(fileBytes, &users)
	if err != nil {
		return nil, errors.Wrap(err, "configuration file structure is invalid")
	}

	return &Config{
		f:     file,
		users: users,
	}, nil
}

func (c *Config) GetUsers() []User {
	return c.users
}

func (c *Config) SaveUserInfo(name string, email string) error {
	//TODO implement me
	panic("implement me")
}

func GetConfigFile() (*os.File, error) {
	pathToFile := GetConfigPath()
	file, err := os.OpenFile(pathToFile, os.O_RDWR, 0777)
	if errors.Is(err, os.ErrNotExist) {
		file, err = os.Create(pathToFile)
	}
	if err != nil {
		return nil, err
	}

	return file, nil
}

func GetConfigPath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s%c%s", dirname, os.PathSeparator, ".gitacc")
}
