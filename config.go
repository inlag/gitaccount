package main

type Config struct {
	Users []User `yaml:"Users"`
}

func NewConfig() (*Config, error) {
	// todo: add handling
	return &Config{}, nil
}

func (c *Config) GetUsers() []User {
	return c.Users
}

func (c *Config) SaveUserInfo(name string, email string) error {
	//TODO implement me
	panic("implement me")
}
