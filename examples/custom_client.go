package examples

import (
	"net/http"

	"github.com/vpofe/go-http-client/gohttp"
)

var (
	customClient = getCustomClient()
)

func getCustomClient() gohttp.Client {
	simpleClient := http.Client{}

	customClient := gohttp.NewBuilder().SetHttpClient(&simpleClient).Build()

	return customClient
}

func MakeRequestToGoogle() {
	customClient.Get("https://api.google.com", nil)
}
