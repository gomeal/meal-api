package themealsdb_client

import (
	"github.com/gomeal/meal-api/internal/clients"
)

type Client struct {
	config clients.TheMealsDbConfig
	cl     clients.HTTPClient
}

func New(config clients.TheMealsDbConfig, httpClient clients.HTTPClient) *Client {
	return &Client{
		config: config,
		cl:     httpClient,
	}
}
