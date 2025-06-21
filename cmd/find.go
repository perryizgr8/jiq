/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/perryizgr8/jiq/jirac"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

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
		fmt.Printf("Showing 1 to %d of %d matching issues\n", len(searchResponse.Issues), searchResponse.Total)
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl := table.New("Key", "Summary", "Status")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
		for _, issue := range searchResponse.Issues {
			tbl.AddRow(issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
		}
		tbl.Print()
		numPrinted := len(searchResponse.Issues)
		for searchResponse.Total > numPrinted {
			fmt.Println("Show more? (Y/n)")
			var cmd string
			fmt.Scanln(&cmd)
			if cmd != "y" && cmd != "Y" && cmd != "" {
				break
			}
			searchResponse, err = jirac.SearchIssues(jc, query, numPrinted)
			if err != nil {
				println("Error searching issues:", err.Error())
				return
			}
			fmt.Printf("Showing %d to %d of %d matching issues\n", numPrinted+1, numPrinted+len(searchResponse.Issues), searchResponse.Total)
			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()
			tbl := table.New("Key", "Summary", "Status")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
			for _, issue := range searchResponse.Issues {
				tbl.AddRow(issue.Key, issue.Fields.Summary, issue.Fields.Status.Name)
			}
			tbl.Print()
			numPrinted += len(searchResponse.Issues)
		}
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
