package examples

import (
	"net/http"
	"testing"

	"github.com/vpofe/go-http-client/gohttp"
)

func TestPostEndpoints(t *testing.T) {
	t.Run("TestErrorPostingToGithub", func(t *testing.T) {
		gohttp.FlushMocks()

		gohttp.AddMock(
			gohttp.Mock{
				Method:       http.MethodPost,
				Url:          "https://api.github.com",
				RequestBody:  `{"name":"testing-repo"}`,
				ResponseBody: `{"status": "ok}`,
			})

		endpoints, err := PostRepoUrl()

		if endpoints == nil {
			t.Error("Endpoints were expected")
		}

		if err != nil {
			t.Error(err.Error())
		}
	})
}
