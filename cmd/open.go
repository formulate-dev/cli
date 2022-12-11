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

func loadConfig() *model.Config {
	config := model.Config{}
	_, err := toml.DecodeFile("formulate.toml", &config)
	errExit(err)
	return &config
}

var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "Open the settings page for this form on formulate.dev",
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()
		url := fmt.Sprintf("https://formulate.dev/manage?id=%s&secret=%s", config.Internal.Id, config.Internal.Secret)
		util.OpenInBrowser(url)
	},
}

var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview unpublished changes to this form on formulate.dev",
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()
		url := fmt.Sprintf("https://formulate.dev/form/%s?preview=1", config.Internal.Id)
		util.OpenInBrowser(url)
	},
}

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "View the last-published version of this form on formulate.dev",
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()
		url := fmt.Sprintf("https://formulate.dev/form/%s", config.Internal.Id)
		util.OpenInBrowser(url)
	},
}

func init() {
	rootCmd.AddCommand(manageCmd)
	rootCmd.AddCommand(previewCmd)
	rootCmd.AddCommand(shareCmd)
}
