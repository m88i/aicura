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
	"encoding/json"
	"net/http"
)

// ServerMessage describes a structure for a bad request error (4.x)
type ServerMessage struct {
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// ClientError describes a base structure for possible errors during Nexus Server REST API calls
type ClientError struct {
	HTTPStatusCode int
	ServerMessages []ServerMessage
	errorMessage   string
}

// Error gets the error message
func (n *ClientError) Error() string {
	return n.errorMessage
}

func newNexusError(resp *http.Response) error {
	nexusError := &ClientError{
		HTTPStatusCode: resp.StatusCode,
		errorMessage:   "Request Failure",
	}
	switch resp.StatusCode {
	case http.StatusBadRequest:
		nexusError.ServerMessages = decodeServerMessage(resp)
	case http.StatusNotFound:
		nexusError.errorMessage = "Not found"
	}
	return nexusError
}

func decodeServerMessage(resp *http.Response) []ServerMessage {
	var messages []ServerMessage
	err := json.NewDecoder(resp.Body).Decode(&messages)
	if err != nil {
		getLogger(false).Warn("Impossible to decode response body into a server message: ", err)
	}
	return messages
}
