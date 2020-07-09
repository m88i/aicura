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
	RawData        interface{}
	RequestBody    interface{}
	errorMessage   string
}

// Error gets the error message
func (n *ClientError) Error() string {
	return n.errorMessage
}

// IsServerError checks if the given error is a server error (500)
func IsServerError(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*ClientError); ok {
		return err.(*ClientError).HTTPStatusCode == http.StatusInternalServerError
	}
	return false
}

// IsNotFound checks if the given error is a Not Found error (404)
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*ClientError); ok {
		return err.(*ClientError).HTTPStatusCode == http.StatusNotFound
	}
	return false
}

// IsAuthenticationError checks if the given error is related to authentication problems (403 or 401)
func IsAuthenticationError(err error) bool {
	if err == nil {
		return false
	}
	if _, ok := err.(*ClientError); ok {
		return err.(*ClientError).HTTPStatusCode == http.StatusUnauthorized ||
			err.(*ClientError).HTTPStatusCode == http.StatusForbidden
	}
	return false
}

func newNexusError(resp *http.Response) error {
	req, res := unmarshalRawData(resp)
	nexusError := &ClientError{
		HTTPStatusCode: resp.StatusCode,
		RawData:        res,
		RequestBody:    req,
	}
	if resp.StatusCode == http.StatusBadRequest {
		nexusError.ServerMessages = decodeServerMessage(resp)
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

func unmarshalRawData(resp *http.Response) (response, request interface{}) {
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		getLogger(false).Warn("Impossible to decode response body into raw data: ", err)
	}
	if resp.Request.Body != nil { // getting the body without this check would panic for empty body requests
		reader, _ := resp.Request.GetBody()
		err = json.NewDecoder(reader).Decode(&request)
		if err != nil {
			getLogger(false).Warn("Impossible to decode request body into raw data: ", err)
		}
	}
	return
}
