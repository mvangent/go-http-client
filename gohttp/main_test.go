package gohttp

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting tests for package gohttp")

	os.Exit(m.Run())
}
