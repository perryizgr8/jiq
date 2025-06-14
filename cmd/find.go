/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/perryizgr8/jiq/jirac"
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
		issues, err := jirac.SearchIssues(jc, query)
		if err != nil {
			println("Error searching issues:", err.Error())
			return
		}
		if len(issues) == 0 {
			println("No issues found for query:", query)
			return
		}
		println("Found", len(issues), "issues for query:", query)
		// for _, issue := range issues {
		// 	println("Issue Key:", issue.Key, "Summary:", issue.Fields.Summary)
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
