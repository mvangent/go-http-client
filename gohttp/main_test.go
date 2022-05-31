package gohttp

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting tests for package gohttp")

	StopMockServer()

	os.Exit(m.Run())
}
