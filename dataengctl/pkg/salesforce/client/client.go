package client

import (
	"github.com/simpleforce/simpleforce"
	"github.com/Mirantis/dataeng/pkg/config")

type Client struct {
	Config *config.SalesForceConfig
}

func (c *Client) Query(q string) (*simpleforce.QueryResult, error) {
	client := simpleforce.NewClient(c.Config.URL, c.Config.ClientID, c.Config.APIVersion)
	err := client.LoginPassword(c.Config.Username, c.Config.Password, c.Config.Token)
	if err != nil {
		return nil, err
	}

	// q := "SELECT+Environment2__r.Name+FROM+Case"
	return client.Query(q) // Note: for Tooling API, use client.Tooling().Query(q)
}


