package jirac

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

func GetIssue(client *jira.Client, issueKey string) (*jira.Issue, error) {
	// Get the issue by its key
	issue, resp, err := client.Issue.Get(issueKey, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting issue %s: %w", issueKey, err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response status for issue %s: %s", issueKey, resp.Status)
	}

	return issue, nil
}
