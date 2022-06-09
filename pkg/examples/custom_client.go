package examples

import (
	"net/http"

	"github.com/vpofe/go-http-client/pkg/gohttp"
)

var (
	customClient = getCustomClient()
)

func getCustomClient() gohttp.Client {
	simpleClient := http.Client{}

	customClient := gohttp.NewBuilder().SetHttpClient(&simpleClient).Build()

	return customClient
}