package examples

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/vpofe/go-http-client/gohttp"
)

func TestGetEndpoints(t *testing.T) {
	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		mock := gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timout getting github endpoints"),
		}

		endpoints, err := GetGithubEndpoints()

	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		}

		endpoints, err := GetGithubEndpoints()

	})

	t.Run("TestNoError", func(t *testing.T) {
		mock := gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		}

		endpoints, err := GetGithubEndpoints()

	})

	if err != nil {
		t.Fail()
	}

	fmt.Println(endpoints.RepositoryUrl)
}
