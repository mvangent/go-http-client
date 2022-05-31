package examples

import (
	"fmt"
	"net/http"
	"time"
)

// type User struct {
// 	FirstName string `json "firstName"`
// 	LastName  string `json "lastName"`
// }

func basicExample() {

	headers := make(http.Header)
	headers.Set("X-Weather", "Sunny in Berlin")

	response, err := httpClient.Get("https://api.github.com", headers)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
}

func main() {

	basicExample()

	time.Sleep(1 * time.Second)

	for i := 0; i < 5; i++ {
		go func() {
			basicExample()
		}()
	}

	time.Sleep(20 * time.Second)
}
