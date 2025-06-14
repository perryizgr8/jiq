package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseURL  string `yaml:"base_url"`
	Username string `yaml:"username"`
	APIToken string `yaml:"api_token"`
}

var config Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jiq",
	Short: "A CLI tool for Jira",
	Long:  `jiq is a command-line interface tool designed to interact with Jira.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func checkAndLoadConfig() {
	configPath := os.ExpandEnv("$HOME/.jiq.yaml")
	_, err := os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			println("Configuration file not found at", configPath)
			println("Please run 'jiq configure' to set up your Jira configuration.")
			os.Exit(1)
		} else {
			println("Error checking configuration file:", err.Error())
			os.Exit(1)
		}
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		println("Error reading configuration file:", err.Error())
		os.Exit(1)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		println("Error parsing configuration file:", err.Error())
		os.Exit(1)
	}
	if config.BaseURL == "" || config.Username == "" || config.APIToken == "" {
		println("Configuration is incomplete. Please run 'jiq configure' to set up your Jira configuration.")
		os.Exit(1)
	}
}
