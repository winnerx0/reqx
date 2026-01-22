package command

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/winnerx0/reqx/internal/utils"
)

var (
	path   string
	silent bool
)

var SendCmd = &cobra.Command{
	Use:   "http [name]",
	Short: "Send HTTP requests ",
	Args:  cobra.ExactArgs(1),
	Long:  "This is used to send the actual http requests using reqx",
	RunE: func(cmd *cobra.Command, args []string) error {

		file := "reqx.yaml"

		var name string

		if len(args) > 0 {
			name = args[0]
		}

		pwd, err := os.Getwd()

		if err != nil {
			return err
		}

		file = filepath.Join(pwd, "reqx.yaml")

		if path != "" {
			file = path
		}

		config, err := utils.Parse(file)

		if err != nil {
			return err
		}

		var namedRequest *utils.Request
		for _, request := range config.Requests {

			if name != "" {

				if request.Name == name {
					namedRequest = &request
				}

			} else {
				err = utils.SendRequest(request, silent, name)

				if err != nil {
					return err
				}
			}

		}

		if name != "" {
			if namedRequest != nil {
				err = utils.SendRequest(*namedRequest, silent, name)
			} else {
				return errors.New("request does not exist in request YAML file")

			}
		}

		if err != nil {
			return err
		}

		return err
	},
	SilenceErrors: true,
	SilenceUsage: true,
}

func init() {
	SendCmd.Flags().StringVarP(&path, "path", "p", "reqx.yaml", "request YAML file path")
	SendCmd.Flags().BoolVarP(&silent, "silent", "s", false, "show only requst status code")
}
