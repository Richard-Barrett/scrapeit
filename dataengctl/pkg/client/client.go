package client

import (
	"fmt"

	dataconfig "github.com/Mirantis/dataeng/pkg/config"
	jira "github.com/Mirantis/dataeng/pkg/jira/client"
	salesforce "github.com/Mirantis/dataeng/pkg/salesforce/client"
)

func NewDataClient(configPath string) *DataClient {
	return &DataClient{
		ConfigFilePath: configPath,
	}
}

type DataClient struct {
	jiraClient       *jira.Client
	salesForceClient *salesforce.Client

	ConfigFilePath string
}

func (dc *DataClient) SalesForceClient() (*salesforce.Client, error) {

	if dc.salesForceClient != nil {
		return dc.salesForceClient, nil
	}

	dataConf := &dataconfig.DataConfig{}
	err := dataconfig.ReadConfig(dc.ConfigFilePath, dataConf)
	if err != nil {
		return nil, err
	}

	if dataConf.SalesforceConfig.Token == "" {
		return nil, fmt.Errorf("You salesforceConfig config at '%s' is missing a token field", dc.ConfigFilePath)
	}

	if dataConf.SalesforceConfig.URL == "" {
		return nil, fmt.Errorf("You salesforceConfig config at '%s' is missing a URL field", dc.ConfigFilePath)
	}

	if dataConf.JiraConfig.Username == "" {
		return nil, fmt.Errorf("You salesforceConfig config at '%s' is missing a username field", dc.ConfigFilePath)
	}

	dc.salesForceClient = &salesforce.Client{
		Config: dataConf.SalesforceConfig,
	}

	return dc.salesForceClient, nil
}

func (dc *DataClient) JiraClient() (*jira.Client, error) {
	if dc.jiraClient != nil {
		return dc.jiraClient, nil
	}

	dataConf := &dataconfig.DataConfig{
		JiraConfig:       &dataconfig.JiraConfig{},
		SalesforceConfig: &dataconfig.SalesForceConfig{},
	}
	err := dataconfig.ReadConfig(dc.ConfigFilePath, dataConf)
	if err != nil {
		return nil, err
	}

	if err = dataConf.JiraConfig.Validate(); err != nil {
		return nil, fmt.Errorf("received an error reading from jira config file at '%s', error is \n '%s'",
			dc.ConfigFilePath,
			err.Error())
	}

	dc.jiraClient = &jira.Client{
		Config: dataConf.JiraConfig,
	}

	return dc.jiraClient, nil
}
