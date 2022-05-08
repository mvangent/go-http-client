package main

import (
	"fmt"
	"io/ioutil"

	"github.com/vpofe/go-http-client/gohttp"
)

func basicExample() {
	client := gohttp.New()

	response, err := client.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func main() {
	basicExample()
}
