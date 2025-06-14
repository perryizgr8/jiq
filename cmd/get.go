package cmd

import (
	"fmt"

	"github.com/perryizgr8/jiq/jirac"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get details of a Jira issue",
	Long:  `Get details of a Jira issue by its key. This command retrieves the issue details from Jira and displays them in a readable format.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkAndLoadConfig()
		if len(args) != 1 {
			println("Please provide an issue key.")
			return
		}
		issueKey := args[0]
		if issueKey == "" {
			println("Please provide an issue key.")
			return
		}
		fmt.Printf("baseURL: %s, username: %s, apiToken: %s\n", config.BaseURL, config.Username, config.APIToken)
		jc, err := jirac.NewClient(config.BaseURL, config.Username, config.APIToken)
		if err != nil {
			println("Error creating Jira client:", err.Error())
			return
		}
		issue, err := jirac.GetIssue(jc, issueKey)
		if err != nil {
			println("Error getting issue:", err.Error())
			return
		}
		println("Issue Key:", issue.Key)
		println("Summary:", issue.Fields.Summary)
		println("Description:", issue.Fields.Description)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
