/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/perryizgr8/jiq/jirac"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type model struct {
	searchResponse jirac.SearchResponse
	cursor         int
}

func initialModel(searchResponse jirac.SearchResponse) model {
	return model{
		searchResponse: searchResponse,
		cursor:         0,
	}
}

func (m model) Init() tea.Cmd {
	// No initialization needed for this model
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.searchResponse.Issues)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			selectedIssue := m.searchResponse.Issues[m.cursor]
			jc, err := jirac.NewClient(config.BaseURL, config.Username, config.APIToken)
			if err != nil {
				println("Error creating Jira client:", err.Error())
				return m, tea.Quit
			}
			issue, err := jirac.GetIssue(jc, selectedIssue.Key)
			if err != nil {
				println("Error getting issue:", err.Error())
				return m, tea.Quit
			}
			fmt.Printf("%s: %s - %s\n", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
			fmt.Println(issue.Fields.Description)
			return m, tea.Quit
		}
	case "m":
		jc, err := jirac.NewClient(config.BaseURL, config.Username, config.APIToken)
		if err != nil {
			println("Error creating Jira client:", err.Error())
			return m, tea.Quit
		}
		if len(m.searchResponse.Issues) == 0 || m.cursor >= len(m.searchResponse.Issues)-1 {
			jc, err := jirac.NewClient(config.BaseURL, config.Username, config.APIToken)
			if err != nil {
				println("Error creating Jira client:", err.Error())
				return m, tea.Quit
			}
		}
		return m, nil
	}
}

func (m model) View() string {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("", "Key", "Summary", "Status")
	var buffer bytes.Buffer
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithWriter(&buffer)

	for i, issue := range m.searchResponse.Issues {
		if i == m.cursor {
			tbl.AddRow(">", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
		} else {
			tbl.AddRow("", issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
		}
	}

	tbl.Print()
	return buffer.String() + fmt.Sprintf("\n%d matching issues", m.searchResponse.Total) +
		"\nUse 'enter' to select, 'm' to load more, 'q' to quit."
}

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find issues in Jira",
	Long:  `Find issues in Jira using fuzzy matching.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkAndLoadConfig()
		if len(args) != 1 {
			println("Please provide a search query.")
			return
		}
		query := args[0]
		if query == "" {
			println("Please provide a search query.")
			return
		}
		jc, err := jirac.NewClient(config.BaseURL, config.Username, config.APIToken)
		if err != nil {
			println("Error creating Jira client:", err.Error())
			return
		}
		searchResponse, err := jirac.SearchIssues(jc, query, 0)
		if err != nil {
			println("Error searching issues:", err.Error())
			return
		}
		p := tea.NewProgram(initialModel(searchResponse))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error starting program: %v\n", err)
			return
		}

		// fmt.Printf("Showing 1 to %d of %d matching issues\n", len(searchResponse.Issues), searchResponse.Total)
		// headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		// columnFmt := color.New(color.FgYellow).SprintfFunc()
		// tbl := table.New("Key", "Summary", "Status")
		// tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		// for _, issue := range searchResponse.Issues {
		// 	tbl.AddRow(issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
		// }
		// tbl.Print()
		// numPrinted := len(searchResponse.Issues)
		// for searchResponse.Total > numPrinted {
		// 	fmt.Println("Show more? (Y/n)")
		// 	var cmd string
		// 	fmt.Scanln(&cmd)
		// 	if cmd != "y" && cmd != "Y" && cmd != "" {
		// 		break
		// 	}
		// 	searchResponse, err = jirac.SearchIssues(jc, query, numPrinted)
		// 	if err != nil {
		// 		println("Error searching issues:", err.Error())
		// 		return
		// 	}
		// 	fmt.Printf("Showing %d to %d of %d matching issues\n", numPrinted+1, numPrinted+len(searchResponse.Issues), searchResponse.Total)
		// 	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		// 	columnFmt := color.New(color.FgYellow).SprintfFunc()
		// 	tbl := table.New("Key", "Summary", "Status")
		// 	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		// 	for _, issue := range searchResponse.Issues {
		// 		tbl.AddRow(issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
		// 	}
		// 	tbl.Print()
		// 	numPrinted += len(searchResponse.Issues)
		// }
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
