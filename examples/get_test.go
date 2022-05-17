package examples

import (
	"fmt"
	"testing"
)

func TestGetEndpoints(t *testing.T) {

	endpoints, err := GetGithubEndpoints()

	if err != nil {
		t.Fail()
	}

	fmt.Println(endpoints.RepositoryUrl)
}
