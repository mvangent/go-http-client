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
		gohttp.FlushMocks()

		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		endpoints, err := GetGithubEndpoints()

		if endpoints != nil {
			t.Error("No endpoints were expected")
		}

		if err == nil {
			t.Error("An error was expected")
		}

		if err.Error() != "timeout getting github endpoints" {
			t.Error(fmt.Printf("Wrong error message, expected: timeout getting github endpoints, and got instead: %s", err.Error()))
		}
	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		gohttp.FlushMocks()

		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})

		endpoints, err := GetGithubEndpoints()

		if endpoints != nil {
			t.Error("No endpoints were expected")
		}

		if err == nil {
			t.Error("An error was expected")
		}

		expectedError := "json: cannot unmarshal number into Go struct field Endpoints.current_user_url of type string"

		if err.Error() != expectedError {
			t.Error(fmt.Printf("Wrong error message, expected: %s, and got instead: %s", expectedError, err.Error()))
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp.FlushMocks()

		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})

		endpoints, err := GetGithubEndpoints()

		if endpoints == nil {
			t.Error("Endpoints were expected")
		}

		if err != nil {
			t.Error("No error was expected")
		}
	})

}
