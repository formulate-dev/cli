/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/formulate-dev/cli/api"
	"github.com/formulate-dev/cli/model"
	"github.com/spf13/cobra"
)

// customIdCmd represents the customId command
var customIdCmd = &cobra.Command{
	Use:   "custom-id [flags] NEW_ID",
	Args:  cobra.ExactArgs(1),
	Short: "Give this form a custom ID",
	Long:  `This lets respondents access the form using a shorter URL: https://formulate.dev/f/<custom_id>`,
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()

		form := model.Form{
			Id:     config.Internal.Id,
			Secret: config.Internal.Secret,
		}

		err := api.SetCustomId(&form, args[0])
		errExit(err)

		config.Internal.CustomId = args[0]
		f, err := os.Create("formulate.toml")
		errExit(err)
		defer f.Close()

		encoder := toml.NewEncoder(f)
		err = encoder.Encode(config)
		errExit(err)

		color.Green("Form '%s' is now accessible at https://formulate.dev/f/%s", config.Title, args[0])
	},
}

func init() {
	rootCmd.AddCommand(customIdCmd)
}
