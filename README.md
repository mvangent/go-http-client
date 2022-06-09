![Coverage](https://img.shields.io/badge/Coverage-89.2%25-brightgreen)
![CI main](https://github.com/vpofe/go-http-client/actions/workflows/go.yml/badge.svg?branch=main)

# go-http-client

A fast, reliable and easily configurable http library around the native http client, with mocking functionality out of the box

## Quick start

File to initialize the http library as a singleton

```
    package examples

    import (
        "net/http"
        "time"

        "github.com/vpofe/go-http-client/gohttp"
        "github.com/vpofe/go-http-client/gomime"
    )

    var (
        httpClient = getHttpClient()
    )

    func getHttpClient() gohttp.Client {
        headers := make(http.Header)
        headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

        client := gohttp.NewBuilder().
            SetHeaders(headers).
            SetUserAgent("vpofe-machine").
            SetConnectionTimeout(2 * time.Second).
            SetResponseTimeout(3 * time.Second).
            Build()

        return client
    }

```

Get example with the singleton instance

```
    package examples

    type Endpoints struct {
        CurrentUserUrl    string `json:"current_user_url"`
        AuthorizationsUrl string `json:"authorizations_url"`
        RepositoryUrl     string `json:"repository_url"`
    }

    func GetGithubEndpoints() (*Endpoints, error) {
        response, err := httpClient.Get("https://api.github.com")

        if err != nil {
            return nil, err
        }

        var endpoints Endpoints

        if err := response.UnmarshalJson(&endpoints); err != nil {
            return nil, err
        }

        return &endpoints, nil
    }
```

## Mocking Requests

Example based on the Get request from above

```
    package examples

    import (
        "errors"
        "fmt"
        "net/http"
        "os"
        "testing"

        "github.com/vpofe/go-http-client/gohttp_mock"
    )

    func TestMain(m *testing.M) {
        fmt.Println("Starting tests for package examples")

        gohttp_mock.MockupServer.Start()

        os.Exit(m.Run())
    }

    func TestGetEndpoints(t *testing.T) {
        t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
            gohttp_mock.MockupServer.Flush()

            gohttp_mock.MockupServer.AddMock(gohttp_mock.Mock{
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
    }
```

### Kudos

The foundation of this library is based on the work by @federicoleon, taught in his Udemy Course: How to develop a productive HTTP client in Golang (Go)
