package examples

import "github.com/vpofe/go-http-client/core"

type RepoUrl struct {
	Name string `json:"name"`
}

func PostRepoUrl() (*core.Response, error) {
	repo := RepoUrl{
		Name: "testing-repo",
	}

	response, err := httpClient.Post("https://api.github.com", repo)

	if err != nil {
		return nil, err
	}

	return response, nil
}
