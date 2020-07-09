//     Copyright 2020 Aicura Nexus Client and/or its authors
//
//     This file is part of Aicura Nexus Client.
//
//     Aicura Nexus Client is free software: you can redistribute it and/or modify
//     it under the terms of the GNU Lesser General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     Aicura Nexus Client is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU Lesser General Public License for more details.
//
//     You should have received a copy of the GNU Lesser General Public License
//     along with Aicura Nexus Client.  If not, see <https://www.gnu.org/licenses/>.

package nexus

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nexusTestBaseURLEnv  = "NEXUS_TEST_BASE_URL"
	nexusTestUserEnv     = "NEXUS_TEST_USER"
	nexusTestPasswordEnv = "NEXUS_TEST_PASSWORD"
)

type serverWrapperBuilder struct {
	responseBody string
	statusCode   int
	t            *testing.T
}

func (m *serverWrapperBuilder) WithResponse(response string) *serverWrapperBuilder {
	m.responseBody = response
	return m
}

func (m *serverWrapperBuilder) WithStatusCode(statusCode int) *serverWrapperBuilder {
	m.statusCode = statusCode
	return m
}

func (m *serverWrapperBuilder) Build() *serverWrapper {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, p, _ := r.BasicAuth()
		assert.Equal(m.t, getUserForTest(), u)
		assert.Equal(m.t, getPasswordForTest(), p)
		if m.statusCode > 0 {
			w.WriteHeader(m.statusCode)
		}
		if len(m.responseBody) > 0 {
			_, err := w.Write([]byte(m.responseBody))
			assert.NoError(m.t, err)
		}
	})
	s := httptest.NewServer(handler)
	mockServer := &serverWrapper{
		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
					return net.Dial(network, s.Listener.Addr().String())
				},
			},
		},
		teardown: s.Close,
	}

	return mockServer
}

type serverWrapper struct {
	httpClient *http.Client
	teardown   func()
}

// Client builds a client reference to be used on unit and integration tests
func (m *serverWrapper) Client() *Client {
	// we are not testing agains a mocked server
	if isRemoteTesting() {
		baseURL := os.Getenv(nexusTestBaseURLEnv)
		getLogger(true).Infof("Testing with baseURL %s", baseURL)
		return NewClient(baseURL).
			WithCredentials(getUserForTest(), getPasswordForTest()).
			Verbose().
			Build()
	}
	return m.MockClient()
}

// moMockClient builds a client reference to use only on unit tests
func (m *serverWrapper) MockClient() *Client {
	return NewClient("http://nexus.com/").
		WithCredentials(getUserForTest(), getPasswordForTest()).
		WithHTTPClient(m.httpClient).
		Verbose().
		Build()
}

// newServerWrapper creates an httptest.Server reference used by unit and integration tests
func newServerWrapper(t *testing.T) *serverWrapperBuilder {
	return &serverWrapperBuilder{t: t}
}

func getUserForTest() string {
	user := os.Getenv(nexusTestUserEnv)
	if len(user) > 0 {
		return user
	}
	return "admin"
}

func getPasswordForTest() string {
	password := os.Getenv(nexusTestPasswordEnv)
	if len(password) > 0 {
		return password
	}
	return "admin123"
}

func isRemoteTesting() bool {
	return len(os.Getenv(nexusTestBaseURLEnv)) > 0
}

// assertRemote will only run the given assert function when testing on remote environments, meaning against an actual Nexus Server
func assertRemote(t *testing.T, assertF func() error) {
	if isRemoteTesting() {
		err := assertF()
		assert.NoError(t, err)
	}
}
