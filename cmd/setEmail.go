package cmd

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/formulate-dev/cli/api"
	"github.com/formulate-dev/cli/model"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var setEmailCmd = &cobra.Command{
	Use:   "set-email [flags] EMAIL_ADDRESS",
	Args:  cobra.ExactArgs(1),
	Short: "Link a form with your email address to receive submissions live",
	Run: func(cmd *cobra.Command, args []string) {
		config := loadConfig()

		form := model.Form{
			Id:     config.Internal.Id,
			Secret: config.Internal.Secret,
		}

		fmt.Printf("Sending verification email to %s...", args[0])
		err := api.VerifyEmail(&form, args[0])
		errExit(err)
		fmt.Println("sent!")

		prompt := promptui.Prompt{Label: fmt.Sprintf("Enter the code we sent to %s", args[0])}
		code, err := prompt.Run()
		errExit(err)

		err = api.SetEmail(&form, code)
		errExit(err)

		config.Internal.Email = args[0]
		f, err := os.Create("formulate.toml")
		errExit(err)
		defer f.Close()

		encoder := toml.NewEncoder(f)
		err = encoder.Encode(config)
		errExit(err)

		color.Green("Form '%s' is now linked with the email %s", config.Title, args[0])
	},
}

func init() {
	rootCmd.AddCommand(setEmailCmd)
}
