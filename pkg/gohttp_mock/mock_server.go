package gohttp_mock

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/vpofe/go-http-client/pkg/core"
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock

	httpClient core.HttpClient
}

var (
	MockupServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

func (m *mockServer) GetClient() core.HttpClient {
	return m.httpClient
}

func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mocks[key] = &mock
}

func (m *mockServer) Flush() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()
	m.mocks = make(map[string]*Mock)
}

func (m *mockServer) IsEnabled() bool {
	return m.enabled
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
