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
)

const (
	nexusTestBaseURLEnv  = "NEXUS_TEST_BASE_URL"
	nexusTestUserEnv     = "NEXUS_TEST_USER"
	nexusTestPasswordEnv = "NEXUS_TEST_PASSWORD"
)

// newMockServer creates an httptest.Server reference used by unit tests
func newMockServer(handler http.Handler) (client *http.Client, teardown func()) {
	s := httptest.NewServer(handler)

	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
	}

	return cli, s.Close
}

// newClientForTest creates a new Client for the test in the given context.
func newClientForTest(client *http.Client) *Client {
	baseURL := os.Getenv(nexusTestBaseURLEnv)
	if len(baseURL) > 0 {
		// we are not testing agains a mocked server
		getLogger(true).Infof("Testing with baseURL %s", baseURL)
		return NewClient(baseURL).
			WithCredentials(getUserForTest(), getPasswordForTest()).
			Verbose().
			Build()
	}
	return NewClient("http://nexus.com/").
		WithCredentials(getUserForTest(), getPasswordForTest()).
		WithHTTPClient(client).
		Verbose().
		Build()

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
