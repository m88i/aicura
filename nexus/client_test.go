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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_IsNonScriptOperationsEnabled_OldServers(t *testing.T) {
	s := newServerWrapper(t).WithStatusCode(http.StatusNotFound).Build()
	defer s.teardown()
	enabled, err := s.MockClient().IsNonScriptOperationsEnabled()
	assert.NoError(t, err)
	assert.False(t, enabled)
}

func TestClient_IsNonScriptOperationsEnabled_NewServers(t *testing.T) {
	s := newServerWrapper(t).WithResponse(listUsersExpected).Build()
	defer s.teardown()
	enabled, err := s.Client().IsNonScriptOperationsEnabled()
	assert.NoError(t, err)
	assert.True(t, enabled)
}
