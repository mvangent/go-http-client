package gohttp

import (
	"errors"
	"fmt"
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
	return method + url + body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	if !m.enabled {
		return nil
	}

	if mock := mockupServer.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: errors.New(fmt.Sprintln("There is no mock matching %s from %s with given body", method, url)),
	}
}
