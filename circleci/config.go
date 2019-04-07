package circleci

import (
	circleci "./circleci-go"
)

type Config struct {
	AuthToken string
}

func (c *Config) NewClient() *circleci.Client {
	return circleci.NewClient(c.AuthToken)
}
