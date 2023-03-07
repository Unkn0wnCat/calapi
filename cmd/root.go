package cmd

import (
	"github.com/Unkn0wnCat/calapi/internal/logger"
	"github.com/spf13/viper"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calapi",
	Short: "Your headless calendar.",
	Long:  `CalAPI allows you to have a headless calendar which can be manipulated using a GraphQL-API.`,
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
	viper.SetEnvPrefix("calapi")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	viper.SetDefault("development", false)
	viper.SetDefault("data_directory", "./data")
	viper.SetDefault("auth.type", "GHOST")
	viper.SetDefault("auth.secret", "hunter2")
	viper.SetDefault("auth.anonymous_read", true)
	viper.SetDefault("auth.ghost.base_url", "https://demo.ghost.io/ghost")
	//viper.SetDefault("auth.ghost.limit_to_roles", nil)

	viper.AutomaticEnv()

	logger.StartLogger()
}
