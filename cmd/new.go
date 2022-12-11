package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/fatih/color"
	"github.com/formulate-dev/cli/api"
	"github.com/formulate-dev/cli/model"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [flags] DIR",
	Short: "Create a new form",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("please set a `DIR` to create for the new form")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		dir := args[0]

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
		err = encoder.Encode(model.Config{
			Title: title,
			Internal: model.ConfigInternal{
				Id:     form.Id,
				Secret: form.Secret,
			},
		})
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

		dirAbsolute, err := filepath.Abs(dir)
		errExit(err)

		color.Green("Created form '%s' at %s.", title, dirAbsolute)
		fmt.Println("\nGetting started:\n- Make changes to `index.js` and run `formulate save` to persist your changes to formulate.dev\n- Run `formulate publish` when you're ready to share your form\n- Docs: https://formulate.dev/docs")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringP("title", "t", "Untitled Form", "What should this form be called?")
	// newCmd.Flags().BoolP("git", "g", false, "Initialize a git repository")
}
