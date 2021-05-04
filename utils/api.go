package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gofrs/uuid"
)

const baseURL = "https://q.trap.jp/api/v3"

type UserMe struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func GetUserMe(token string) (*UserMe, error) {
	bytes, err := APIGetRequest(token, "/users/me")
	if err != nil {
		return nil, err
	}

	userMe := new(UserMe)
	if err := json.Unmarshal(bytes, userMe); err != nil {
		return nil, err
	}
	return userMe, nil
}

func APIGetRequest(token, endpoint string) ([]byte, error) {
	if token == "" {
		return nil, errors.New(http.StatusText(http.StatusUnauthorized))
	}
	req, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		return nil, errors.New(http.StatusText(res.StatusCode))
	}
	return ioutil.ReadAll(res.Body)
}
