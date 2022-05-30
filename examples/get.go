package examples

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

	var endpoints Endpoints

	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}

	return &endpoints, nil
}
