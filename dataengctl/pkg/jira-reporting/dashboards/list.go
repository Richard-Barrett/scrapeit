package dashboard

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Mirantis/dataeng/pkg/client"
)


func List(dataClient *client.DataClient) error {
	jiraClient, err := dataClient.JiraClient()
	if err!= nil {
		return err
	}

	url := strings.Join([]string{jiraClient.Config.URL, "rest/api/latest/dashboard"}, "/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(jiraClient.Config.Username, jiraClient.Config.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}