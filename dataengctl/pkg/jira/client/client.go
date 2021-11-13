package client

import (
	"context"
	"fmt"

	"github.com/Mirantis/dataeng/pkg/config"
	"github.com/andygrunwald/go-jira"
)

type Client struct {
	Config *config.JiraConfig
	client *jira.Client
}

func (c *Client) IssueIterator(opts IssueIteratorOptions) (*IssueIterator, error) {

	tp := jira.BasicAuthTransport{
		Username: c.Config.Username,
		Password: c.Config.Token,
	}
	jiraClient, err := jira.NewClient(tp.Client(), c.Config.URL)
	if err != nil {
		return nil, err
	}
	c.client = jiraClient

	return NewIssueIterator(context.Background(), jiraClient, opts)
}

type IssueIterator struct {
	client       *jira.Client
	query        string
	expand       string
	issues       []jira.Issue
	pageSize     int
	startAt      int
	endAt        int
	totalResults int
	index        int
}

type IssueIteratorOptions struct {
	Query    string
	Expand   string
	PageSize int
}

func NewIssueIterator(ctx context.Context, c *jira.Client, opts IssueIteratorOptions) (*IssueIterator, error) {

	if opts.Query == "" {
		return nil, fmt.Errorf("must pass non-nil query")
	}

	it := &IssueIterator{
		client:   c,
		query:    opts.Query,
		expand:   opts.Expand,
		pageSize: opts.PageSize,
	}
	return it, it.more(ctx)
}

func (it *IssueIterator) more(ctx context.Context) error {
	issues, response, err := it.client.Issue.SearchWithContext(
		ctx,
		it.query,
		&jira.SearchOptions{
			StartAt:    it.endAt,
			MaxResults: it.pageSize,
			Expand:     it.expand,
		},
	)
	if err != nil {
		return err
	}

	it.issues = issues
	it.startAt = response.StartAt
	it.totalResults = response.Total
	it.endAt = it.endAt + response.MaxResults
	if it.endAt > response.Total {
		it.endAt = response.Total
	}
	return nil
}

func (it *IssueIterator) Next(ctx context.Context) (*jira.Issue, error) {
	it.index = it.index + 1

	if it.index > it.totalResults {
		return nil, nil
	}

	if it.index > it.endAt {
		if err := it.more(ctx); err != nil {
			return nil, fmt.Errorf("error getting issues: %v\n", err)
		}
	}
	return &it.issues[it.index-it.startAt-1], nil
}
