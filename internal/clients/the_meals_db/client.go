package themealsdb_client

import (
	"time"

	"github.com/gomeal/meal-api/internal/clients"
)

type Config interface {
	Url() string
	Timeout() time.Duration
}

type Client struct {
	config Config
	cl     clients.HTTPClient
}

func New(config Config, httpClient clients.HTTPClient) *Client {
	return &Client{
		config: config,
		cl:     httpClient,
	}
}
