package jirac

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

func NewClient(baseURL, username, apiToken string) (*jira.Client, error) {
	// Create a new Jira client
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: apiToken,
	}

	client, err := jira.NewClient(tp.Client(), baseURL)
	if err != nil {
		return nil, fmt.Errorf("error creating Jira client: %w", err)
	}

	return client, nil
}
