package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var client = &http.Client{}

func SendRequest(request Request, silent bool, name string) error {
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
			return nil
}