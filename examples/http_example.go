package examples

import (
	"github.com/vpofe/go-http-client/gohttp"
)

func basicExample() {
	client := gohttp.New()

	client.Get()
}
