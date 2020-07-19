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

import "encoding/json"

const repositoriesMavenProxyPath = "/repositories/maven/proxy"

// MavenVersionPolicy ...
type MavenVersionPolicy string

const (
	// VersionPolicyRelease ...
	VersionPolicyRelease MavenVersionPolicy = "RELEASE"
	// VersionPolicyMixed ...
	VersionPolicyMixed MavenVersionPolicy = "MIXED"
	// VersionPolicySnapshot ...
	VersionPolicySnapshot MavenVersionPolicy = "SNAPSHOT"
)

// MavenLayoutPolicy ...
type MavenLayoutPolicy string

const (
	// LayoutPolicyStrict ...
	LayoutPolicyStrict MavenLayoutPolicy = "STRICT"
	// LayoutPolicyPermissive ...
	LayoutPolicyPermissive MavenLayoutPolicy = "PERMISSIVE"
)

// Maven ...
type Maven struct {
	VersionPolicy MavenVersionPolicy `json:"versionPolicy,omitempty"`
	LayoutPolicy  MavenLayoutPolicy  `json:"layoutPolicy,omitempty"`
}

// HTTPClient details of how to connect to a Maven Proxy Repository
type HTTPClient struct {
	AutoBlock      bool            `json:"autoBlock,omitempty"`
	Blocked        *bool           `json:"blocked,omitempty"`
	Connection     *Connection     `json:"connection,omitempty"`
	Authentication *Authentication `json:"authentication,omitempty"`
}

// Connection for a HTTPClient
type Connection struct {
	Retries                 *int    `json:"retries,omitempty"`
	UserAgentSuffix         *string `json:"userAgentSuffix,omitempty"`
	Timeout                 *int32  `json:"timeout,omitempty"`
	EnableCircularRedirects bool    `json:"enableCircularRedirects,omitempty"`
	EnableCookies           bool    `json:"enableCookies,omitempty"`
}

// Authentication for a HTTPClient
type Authentication struct {
	Type       string `json:"type,omitempty"`
	Username   string `json:"username,omitempty"`
	NtlmHost   string `json:"ntlmHost,omitempty"`
	NtlmDomain string `json:"ntlmDomain,omitempty"`
}

// MavenProxyRepository basic structure for a Maven Proxy Repository
type MavenProxyRepository struct {
	Repository    `json:",inline"`
	HTTPClient    HTTPClient    `json:"httpClient,omitempty"`
	Maven         Maven         `json:"maven,omitempty"`
	NegativeCache NegativeCache `json:"negativeCache,omitempty"`
	Proxy         Proxy         `json:"proxy,omitempty"`
	Storage       Storage       `json:"storage,omitempty"`
	CleanUp       *CleanUp      `json:"cleanup,omitempty"`
	RoutingRule   string        `json:"routingRule,omitempty"`
}

// MavenGroupRepository structure for Maven Group Repository
type MavenGroupRepository struct {
	Repository `json:",inline"`
	Group      MavenGroup `json:"group"`
	Storage    Storage    `json:"storage,omitempty"`
}

// MavenGroup describes a collection of Maven repositories in a given group
type MavenGroup struct {
	MemberNames []string `json:"memberNames"`
}

type mavenRepositoryService service

func (m *mavenRepositoryService) list(repoType RepositoryType, unmarshalValue interface{}) error {
	req, err := m.client.get(m.client.appendVersion(repositoriesPath), "")
	if err != nil {
		return err
	}
	var jsonRepos json.RawMessage
	_, err = m.client.do(req, &jsonRepos)
	if err != nil {
		return err
	}
	filtered, err := filterRepositoryJSONByFormat(jsonRepos, repositoryFormatMaven2, repoType)
	if err != nil {
		return err
	}
	if filtered != nil {
		err = json.Unmarshal(filtered, unmarshalValue)
		return err
	}
	return nil
}
