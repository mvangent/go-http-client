package gohttp_mock

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"sync"
)

var (
	mockupServer = mockServer{mocks: make(map[string]*Mock)}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func FlushMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()
	mockupServer.mocks = make(map[string]*Mock)
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := sha256.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))

	key := hex.EncodeToString(hasher.Sum(nil))

	return key
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)

	if body == "" {
		return ""
	}

	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, " ", "")

	return body
}

func GetMock(method, url, body string) *Mock {
	if !mockupServer.enabled {
		return nil
	}

	if mock := mockupServer.mocks[mockupServer.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	errorMsg := fmt.Sprintf("There is no mock matching %s from %s with given body \n", method, url)

	return &Mock{
		Error: errors.New(errorMsg),
	}
}
