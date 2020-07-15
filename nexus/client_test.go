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
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_IsNonScriptOperationsEnabled_OldServers(t *testing.T) {
	s := newServerWrapper(t).WithStatusCode(http.StatusNotFound).Build()
	defer s.teardown()
	enabled, err := s.MockClient().ScriptsRequired()
	assert.NoError(t, err)
	assert.True(t, enabled)
}

func TestClient_IsNonScriptOperationsEnabled_NewServers(t *testing.T) {
	s := newServerWrapper(t).WithResponse(listUsersExpected).Build()
	defer s.teardown()
	enabled, err := s.Client().ScriptsRequired()
	assert.NoError(t, err)
	assert.False(t, enabled)
}

func TestClient_put(t *testing.T) {
	c := NewDefaultClient("")
	apiPath := "test-path"
	query := "test-query"
	body := "test-body"
	req, err := c.put(apiPath, query, body)
	assert.Nil(t, err)
	assert.True(t, strings.HasSuffix(req.URL.Path, apiPath))
	assert.Equal(t, query, req.URL.RawQuery)

	reqBody := make([]byte, 2*len(body))
	_, err = req.Body.Read(reqBody)
	assert.Nil(t, err)
	assert.Contains(t, string(reqBody), body)
}

func TestClient_SetCredentials(t *testing.T) {
	c := &Client{}
	username := "test-username"
	password := "test-password"
	c.SetCredentials(username, password)
	assert.Equal(t, username, c.username)
	assert.Equal(t, password, c.password)
}

func TestClient_SetUsername(t *testing.T) {
	c := &Client{}
	username := "test-username"
	c.SetUsername(username)
	assert.Equal(t, username, c.username)
}

func TestClient_SetPassword(t *testing.T) {
	c := &Client{}
	password := "test-password"
	c.SetPassword(password)
	assert.Equal(t, password, c.password)
}
