package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserRequest struct {
	accessToken string
}

type UserResponse struct {
	Id    int64  `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func NewUserRequest(accessToken string) *UserRequest {
	return &UserRequest{
		accessToken: accessToken,
	}
}

func (r *UserRequest) Execute() (*UserResponse, error) {
	response, err := r.request()
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bytes := []byte(body)
	var userResponse UserResponse
	e := json.Unmarshal(bytes, &userResponse)
	if e != nil {
		return nil, e
	}

	return &userResponse, nil
}

func (r *UserRequest) request() (*http.Response, error) {
	request, err := r.createRequest()
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(request)
}

func (r *UserRequest) createRequest() (*http.Request, error) {
	uri := "https://api.github.com/user"
	request, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/vnd.github.v3+json")
	request.Header.Set("Authorization", fmt.Sprintf("token %s", r.accessToken))

	return request, nil
}
