package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpMethod := "GET"
	url := "https://api.github.com"

	client := http.Client{}

	request, err := http.NewRequest(httpMethod, url, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept", "application/xml")

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)
	fmt.Println(string(bytes))
}
