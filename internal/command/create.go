package command

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
	"github.com/winnerx0/reqx/internal/utils"
)

var CreateCmd = &cobra.Command{

	Use:   "create",
	Short: "Create a reqx.yaml file",
	RunE: func(cmd *cobra.Command, args []string) error {

		file, err := os.Create("reqx.yaml")

		if err != nil {
			return err
		}

		requests := utils.Config{Requests: []utils.Request{
		}}

		yamlBytes, err := yaml.Marshal(requests)

		if err != nil {
			return err
		}

		_, err = file.Write(yamlBytes)

		if err != nil {
			return err
		}

		fmt.Println("reqx.yaml file created in current working directory")

		return nil
	},
}
