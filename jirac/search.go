package jirac

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

func SearchIssues(client *jira.Client, query string) ([]jira.Issue, error) {
	searchOptions := &jira.SearchOptions{
		MaxResults: 10,
	}
	jql := fmt.Sprintf("textfields ~ \"%s*\"", query)

	issues, resp, err := client.Issue.Search(jql, searchOptions)
	if err != nil {
		return nil, fmt.Errorf("error searching issues: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return issues, nil
}
