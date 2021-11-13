package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DataConfig struct {
	JiraConfig       *JiraConfig       `yaml:"jiraConfig"`
	SalesforceConfig *SalesForceConfig `yaml:"salesForceConfig"`
}

type JiraConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Token    string `yaml:"token"`
}

type SalesForceConfig struct {
	// TODO populate this correct fields
	URL        string `yaml:"url"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Token      string `yaml:"token"`
	ClientID   string `yaml:"clientID"`
	APIVersion string `yaml:"apiVersion"`
}

func ReadConfig(path string, config *DataConfig) error {
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(cfgFile, config)
	return err
}

func (c *JiraConfig) Validate() error {
	if c.Token == "" {
		return fmt.Errorf("Your jiraConfig is missing a token field")
	}

	if c.URL == "" {
		return fmt.Errorf("Your jiraConfig is missing a URL field")
	}

	if c.Username == "" {
		return fmt.Errorf("Your jiraConfig is missing a username field")
	}
	return nil
}
