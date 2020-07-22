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
	"encoding/json"
	"fmt"

	"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
)

const (
	repositoriesPath = "/repositories"
)

// Repository is the base structure for Nexus repositories
type Repository struct {
	Name   string            `json:"name,omitempty"`
	Format *RepositoryFormat `json:"format,omitempty"`
	Type   *RepositoryType   `json:"type,omitempty"`
	URL    *string           `json:"url,omitempty"`
	Online *bool             `json:"online,omitempty"`
}

// Storage ...
type Storage struct {
	StrictContentTypeValidation bool   `json:"strictContentTypeValidation,omitempty"`
	BlobStoreName               string `json:"blobStoreName,omitempty"`
}

// CleanUp ...
type CleanUp struct {
	PolicyNames []string `json:"policyNames,omitempty"`
}

// Proxy ...
type Proxy struct {
	ContentMaxAge  int32  `json:"contentMaxAge,omitempty"`
	MetadataMaxAge int32  `json:"metadataMaxAge,omitempty"`
	RemoteURL      string `json:"remoteUrl,omitempty"`
}

// NegativeCache ...
type NegativeCache struct {
	Enabled    bool  `json:"enabled,omitempty"`
	TimeToLive int32 `json:"timeToLive,omitempty"`
}

// RepositoryType ...
type RepositoryType string

const (
	// RepositoryTypeHosted ...
	RepositoryTypeHosted RepositoryType = "hosted"
	// RepositoryTypeProxy ...
	RepositoryTypeProxy RepositoryType = "proxy"
	// RepositoryTypeGroup ...
	RepositoryTypeGroup RepositoryType = "group"
)

// RepositoryFormat describes supported API repositories format
type RepositoryFormat string

const (
	// RepositoryFormatMaven2 ...
	RepositoryFormatMaven2 RepositoryFormat = "maven2"
)

func filterRepository(jsonRepos json.RawMessage, format RepositoryFormat, repoType RepositoryType) ([]byte, error) {
	// converts into a generic Go type (the JsonPath parser does not work with raw/bytes/string types)
	jsonData := interface{}(nil)
	if err := json.Unmarshal(jsonRepos, &jsonData); err != nil {
		return nil, err
	}
	builder := gval.Full(jsonpath.PlaceholderExtension())
	// filter using JsonPath only the repositories that we are interested to
	path, err := builder.NewEvaluable(fmt.Sprintf(`$..[?(@.format == "%s" && @.type == "%s")]`, format, repoType))
	if err != nil {
		return nil, err
	}
	res, err := path(context.Background(), jsonData)
	if err != nil {
		return nil, err
	}
	// marshal the result back to json, so it can be converted to any type
	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return b, nil
}
