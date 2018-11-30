package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var configFile string
var export string

var rootCmd = &cobra.Command{
	Use:   "twitz",
	Short: "This is a CLI App to parse, retrieve information about, and manage Twitter User Account information.",
	Long: `This is a CLI App for retrieving, storing, and organizing information about Twitter Accounts.

The CLI has the following commands. parse, config, findem, and webscanem. For more information check out the Github Project
page here https://adron.github.io/twitz/.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	configFile = ".twitz.yaml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil() {
		fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}
}
