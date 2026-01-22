package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/winnerx0/reqx/internal/utils"
)

var client = &http.Client{}

var (
	path   string
	silent bool
)

var SendCmd = &cobra.Command{
	Use:   "http [name]",
	Short: "Send HTTP requests ",
	Args:  cobra.MaximumNArgs(1),
	Long:  "This is used to send the actual http requests using reqx",
	RunE: func(cmd *cobra.Command, args []string) error {

		file := "reqx.yaml"

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

		for _, request := range config.Requests {

			bodyBytes, err := json.Marshal(request.Body)

			if err != nil {
				return err
			}

			jsonData := bytes.NewReader(bodyBytes)

			req, err := http.NewRequest(string(request.Method), request.Url, jsonData)

			if err != nil {
				return err
			}

			for k, v := range request.Headers {

				req.Header.Add(k, v)
			}

			response, err := client.Do(req)

			if err != nil {
				return err
			}

			if silent {
				fmt.Println("Name: " + request.Name + "\t Status: " + response.Status + "\n")
			} else {

				responsebytes, err := io.ReadAll(response.Body)

				if err != nil {
					return err
				}
				fmt.Println("Name: " + request.Name + "\n\n" + string(responsebytes) + "\n")
			}

		}

		return err
	},
}

func init() {
	SendCmd.Flags().StringVarP(&path, "path", "p", "reqx.yaml", "request YAML file path")
	SendCmd.Flags().BoolVarP(&silent, "silent", "s", false, "show only requst status code")
}
