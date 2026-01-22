package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/winnerx0/reqx/internal/utils"
)

var client = &http.Client{}

var SendCmd = &cobra.Command{
	Use:   "http [request.yaml]",
	Short: "Send HTTP requests ",
	Args:  cobra.MaximumNArgs(1),
	Long:  "This is used to send the actual http requests using reqx",
	RunE: func(cmd *cobra.Command, args []string) error {

		file := "request.yaml"

		if len(args) > 0 {
			file = args[0]
		}

		config, err := utils.Parse(file)

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

			responsebytes, err := io.ReadAll(response.Body)

			if err != nil {
				return err
			}

			fmt.Println("Name: " + request.Name + "\n\n" + string(responsebytes) + "\n")

		}

		return err
	},
}
