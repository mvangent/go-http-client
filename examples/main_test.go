package examples

import (
	"fmt"
	"os"
	"testing"

	"github.com/vpofe/go-http-client/gohttp_mock"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting tests for package examples")

	gohttp_mock.StartMockServer()

	os.Exit(m.Run())
}
