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
		return nil, errors.Wrap(err, "reading file is failed")
	}

	var users = make([]User, 0)
	if len(fileBytes) > 0 {
		err = yaml.Unmarshal(fileBytes, &users)
		if err != nil {
			return nil, errors.Wrap(err, "configuration file structure is invalid")
		}
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
	c.users = append(c.users, User{
		Name:  name,
		Email: email,
	})

	usersBytes, err := yaml.Marshal(c.users)
	if err != nil {
		return err
	}

	err = c.saveFile(usersBytes)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) saveFile(content []byte) error {

	_ = c.f.Truncate(0)
	_, _ = c.f.Seek(0, 0)

	_, err := c.f.Write(content)
	if err != nil {
		return err
	}

	err = c.f.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) CloseFile() error {
	return c.f.Close()
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
