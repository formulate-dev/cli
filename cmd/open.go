/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/formulate-dev/cli/model"
	"github.com/formulate-dev/cli/util"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the settings page for this form on formulate.dev",
	Run: func(cmd *cobra.Command, args []string) {
		config := model.Config{}
		_, err := toml.DecodeFile("formulate.toml", &config)
		errExit(err)

		url := fmt.Sprintf("https://formulate.dev/manage?id=%s&secret=%s", config.Id, config.Secret)
		util.OpenInBrowser(url)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
