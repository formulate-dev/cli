/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/formulate-dev/cli/api"
	"github.com/formulate-dev/cli/model"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:     "save",
	Aliases: []string{"s"},
	Short:   "Persist local changes to formulate.dev",
	Run: func(cmd *cobra.Command, args []string) {
		config := model.Config{}
		_, err := toml.DecodeFile("formulate.toml", &config)
		errExit(err)

		script, err := ioutil.ReadFile("index.js")
		errExit(err)

		form := model.Form{
			Id:     config.Id,
			Secret: config.Secret,
			Script: string(script),
			Title:  config.Title,
		}

		err = api.UpdateForm(form)
		errExit(err)

		color.Green("Updated form '%s'", config.Title)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().BoolP("preview", "p", false, "Open this form's Preview URL in a new browser window")
}
