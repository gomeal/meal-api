package themealsdb_client

import (
	"net/http"
	"time"
)

type Config interface {
	Url() string
	Timeout() time.Duration
}

type Client struct {
	config Config
	cl     *http.Client
}

func New(config Config) *Client {
	return &Client{
		config: config,
		cl: &http.Client{
			Timeout: config.Timeout(),
		},
	}
}
