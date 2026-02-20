package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var client = &http.Client{}

func SendRequest(request Request, silent bool, field string, name string) error {
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

	defer response.Body.Close()

	if silent {
		fmt.Printf("Name: %s\t Status: %s", request.Name, response.Status)
	} else if field != "" {

		responsebytes, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}

		var data map[string]any

		err = json.Unmarshal(responsebytes, &data)

		if err != nil {
			return err
		}

		details, ok := GetBodyData(data, field)

		if !ok {
			return errors.New("Field does not exist in response")
		}

		fmt.Printf("Name: %s\n\nValue: %s", request.Name, details)


	} else {

		responsebytes, err := io.ReadAll(response.Body)

		if err != nil {
			return err
		}
		fmt.Printf("Name: %s\n\nValue: %s", request.Name, string(responsebytes))
	}
	return nil
}

func GetBodyData(data map[string]any, path string) (any, bool) {

	keys := strings.Split(path, ".")

	var current any = data

	for _, key := range keys {

		m, ok := current.(map[string]any)

		if !ok {
			return nil, false
		}

		current, ok = m[key]

		if !ok {
			return nil, false
		}
	}

	return current, true
}
