package examples

import "fmt"

type Endpoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetGithubEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Status code: %d \n", response.StatusCode())
	fmt.Printf("Status: %s \n", response.Status())
	fmt.Printf("Body: %s\n", response.String())

	var endpoints Endpoints

	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	fmt.Printf("Repository url: %s \n", endpoints.RepositoryUrl)

	return &endpoints, nil
}
