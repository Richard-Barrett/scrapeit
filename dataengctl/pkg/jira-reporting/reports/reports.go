package reports

import (
	"encoding/json"
	"io"

	"github.com/Mirantis/dataeng/pkg/client"
	"github.com/andygrunwald/go-jira"
)

type Options struct {
	Output     string
	Config     *client.DataClient
}

func (lo *Options) List(w io.Writer) error {
	jiraClient, err := lo.Config.JiraClient()
	if err!= nil {
		return err
	}
	tp := jira.BasicAuthTransport{
		Username: jiraClient.Config.Username,
		Password: jiraClient.Config.Token,
	}

	client, err := jira.NewClient(tp.Client(), jiraClient.Config.URL)
	if err != nil {
		return err
	}
	boardList, _, err := client.Board.GetAllBoards(&jira.BoardListOptions{})
    if err  != nil {
		return err
	}
	b, err := json.Marshal(boardList)
    if err  != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}