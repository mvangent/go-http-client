package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vpofe/go-http-client/gohttp"
)

type User struct {
	FirstName string `json "firstName"`
	LastName  string `json "lastName"`
}

func getGithubClient() gohttp.HttpClient {
	githubHttpClient := gohttp.New()

	commonHeaders := make(http.Header)
	// commonHeaders.Set("Authorization", "Bearer 123456")

	githubHttpClient.SetHeaders(commonHeaders)

	return githubHttpClient
}

func basicExample() {
	client := getGithubClient()

	headers := make(http.Header)
	headers.Set("X-Weather", "Sunny in Berlin")

	response, err := client.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func createUser(user User) {
	client := getGithubClient()

	response, err := client.Post("https://api.github.com", nil, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytes))
}

func main() {
	basicExample()

	//futureMusic := User{"Future", "Music"}
	//createUser(futureMusic)
}
