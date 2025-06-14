/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure jiq",
	Long:  `Configure jiq to work with your Jira instance and workflow.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Jira base URL:")
		var baseURL string
		fmt.Scanln(&baseURL)
		fmt.Print("Jira username:")
		var username string
		fmt.Scanln(&username)
		fmt.Print("Jira API token:")
		var apiToken string
		// Scan without echoing input
		terminal := make([]byte, 1)
		for {
			os.Stdin.Read(terminal)
			if terminal[0] == '\n' {
				break
			}
			apiToken += string(terminal[0])
		}
		fmt.Println("Saving configuration to ~/.jiq.yaml")
		var cfg = map[string]string{
			"base_url":  baseURL,
			"username":  username,
			"api_token": apiToken,
		}
		data, err := yaml.Marshal(cfg)
		if err != nil {
			fmt.Println("Error marshalling configuration:", err)
			return
		}
		err = os.WriteFile(os.ExpandEnv("$HOME/.jiq.yaml"), data, 0644)
		if err != nil {
			fmt.Println("Error writing configuration file:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
