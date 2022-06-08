package examples

import (
	"net/http"
	"time"

	"github.com/vpofe/go-http-client/pkg/gohttp"
	"github.com/vpofe/go-http-client/pkg/gomime"
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
