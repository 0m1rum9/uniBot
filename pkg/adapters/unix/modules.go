package unix

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetModules(token string) ([]Module, error) {
	m := ModulesResponse{} // uniX api strangely returns json as modules: []Modules

	req, err := http.NewRequest("GET",
		"https://uni-x.almv.kz/api/modules",
		nil,
	)
	if err != nil {
		return []Module{}, err
	}
	req.Header.Add(
		"Authorization", fmt.Sprintf("Bearer %s", token),
	)
	req.Header.Add(
		"Content-Type","application/json",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Module{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []Module{}, ErrInvalidToken
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&m)

	return m.Modules, nil
}
