package cmd

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/formulate-dev/cli/api"
	"github.com/formulate-dev/cli/model"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new form",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		dir, _ := cmd.Flags().GetString("dir")

		// 1. Create directory
		err := os.Mkdir(dir, 0755)
		errExit(err)

		// 2. Create form on formulate.dev
		form, err := api.CreateForm(title)
		errExit(err)

		// 3. Create script
		err = ioutil.WriteFile(path.Join(dir, "index.js"), []byte(form.Script), 0755)
		errExit(err)

		// 4. Create config file
		f, err := os.Create(path.Join(dir, "formulate.toml"))
		errExit(err)
		defer f.Close()

		encoder := toml.NewEncoder(f)
		err = encoder.Encode(model.Config{Id: form.Id, Secret: form.Secret})
		errExit(err)

		// 5. Create d.ts
		r, err := http.Get("https://raw.githubusercontent.com/formulate-dev/typings/main/index.d.ts")
		errExit(err)
		defer r.Body.Close()

		f, err = os.Create(path.Join(dir, "index.d.ts"))
		errExit(err)
		defer f.Close()

		_, err = io.Copy(f, r.Body)
		errExit(err)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("title", "t", "Untitled Form", "What should this form be called?")
	newCmd.Flags().StringP("dir", "d", "untitled-form", "Local directory to create")
	// newCmd.Flags().BoolP("git", "g", false, "Initialize a git repository")
}
