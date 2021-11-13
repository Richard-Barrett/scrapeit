package boards

import (
	"encoding/json"
	"io"
	//"log"

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

	//log.Printf("I am trying to list boards")
	client, err := jira.NewClient(tp.Client(), jiraClient.Config.URL)
	if err != nil {
		return err
	}
	//log.Printf("I have built jira client")
	boardList, _, err := client.Board.GetAllBoards(&jira.BoardListOptions{})
    if err  != nil {
		return err
	}
	//log.Printf("I got all the boards")
	b, err := json.Marshal(boardList)
    if err  != nil {
		return err
	}

	//log.Printf("I ve marshaled the boards")

	_, err = w.Write(b)
	//log.Printf("I've written the boards, but maybe i got an error")
	return err
}