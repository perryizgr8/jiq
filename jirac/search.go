package jirac

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
)

type SearchResponse struct {
	Total  int          `json:"total"`
	Issues []jira.Issue `json:"issues"`
}

func SearchIssues(client *jira.Client, query string, startAt int) (SearchResponse, error) {
	searchOptions := &jira.SearchOptions{
		MaxResults: 10,
		StartAt:    startAt,
	}
	jql := fmt.Sprintf("textfields ~ \"%s*\" ORDER BY updated DESC", query)

	issues, resp, err := client.Issue.Search(jql, searchOptions)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("error searching issues: %w", err)
	}

	if resp.StatusCode != 200 {
		return SearchResponse{}, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return SearchResponse{
		Total:  resp.Total,
		Issues: issues,
	}, nil
}
