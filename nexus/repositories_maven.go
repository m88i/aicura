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

	"golang.org/x/sync/errgroup"
)

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

// MavenProxyRepositoryService service to handle all Maven Proxy Repositories operations
type MavenProxyRepositoryService service

// Add adds new Proxy Maven repositories to the Nexus Server
func (m *MavenProxyRepositoryService) Add(repositories ...MavenProxyRepository) error {
	if len(repositories) == 0 {
		m.logger.Warnf("Called AddRepository with no repositories to add")
		return nil
	}
	errs, _ := errgroup.WithContext(context.Background())
	for _, repo := range repositories {
		errs.Go(func() error {
			req, err := m.client.post(m.client.appendVersion(repositoriesMavenProxyPath), "", repo)
			if err != nil {
				return err
			}
			_, err = m.client.do(req, nil)
			return err
		})
	}
	return errs.Wait()
}

// List lists all maven repositories from the Nexus Server
func (m *MavenProxyRepositoryService) List() ([]MavenProxyRepository, error) {
	req, err := m.client.get(m.client.appendVersion(repositoriesPath), "")
	if err != nil {
		return nil, err
	}
	var jsonRepos json.RawMessage
	_, err = m.client.do(req, &jsonRepos)
	if err != nil {
		return nil, err
	}
	filtered, err := filterRepositoryJSONByFormat(jsonRepos, repositoryFormatMaven2)
	if err != nil {
		return nil, err
	}
	if filtered != nil {
		repositories := []MavenProxyRepository{}
		err = json.Unmarshal(filtered, &repositories)
		return repositories, err
	}
	return nil, nil
}

// GetRepoByName gets the repository by name or nil if not found
func (m *MavenProxyRepositoryService) GetRepoByName(name string) (*MavenProxyRepository, error) {
	repos, err := m.List()
	if err != nil {
		return nil, err
	}
	for _, repo := range repos {
		if repo.Name == name {
			return &repo, nil
		}
	}
	return nil, nil
}
