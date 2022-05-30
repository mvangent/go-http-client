package examples

import "github.com/vpofe/go-http-client/gohttp"

type RepoUrl struct {
	Name string `json:"name"`
}

func PostRepoUrl() (*gohttp.Response, error) {
	repo := RepoUrl{
		Name: "testing-repo",
	}

	response, err := httpClient.Post("https://api.github.com", nil, repo)

	if err != nil {
		return nil, err
	}

	return response, nil
}
