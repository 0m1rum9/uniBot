// Package unix adapts api of https://uni-x.almv.kz/api
package unix

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Login(login, password string) (token string, err error) {
	jsonData, err := json.Marshal(map[string]string{
		"login":    login,
		"password": password,
	})

	if err != nil {
		return "", err
	}

	resp, err := http.Post(
		"https://uni-x.almv.kz/api/auth/login",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusCreated {
		return "", ErrInvalidCredentials
	}

	defer resp.Body.Close()
	str, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	response := make(map[string]any)
	err = json.Unmarshal([]byte(string(str)), &response)
	if err != nil {
		return "", err
	}
	token = response["token"].(string)

	return token, nil
}
