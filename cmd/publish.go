/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/formulate-dev/cli/api"
	"github.com/formulate-dev/cli/model"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Persist local changes to formulate.dev and publish them",
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()

		script, err := ioutil.ReadFile("index.js")
		errExit(err)

		form := model.Form{
			Id:     config.Id,
			Secret: config.Secret,
			Script: string(script),
			Title:  config.Title,
		}

		err = api.PublishForm(form)
		errExit(err)

		color.Green("Published form '%s'.", config.Title)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
}
