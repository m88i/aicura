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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMavenProxyRepositoryService_List(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositoriesMockData).Build()
	defer s.teardown()
	repos, err := s.Client().MavenProxyRepositoryService.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, repos)
	for _, repo := range repos {
		if repo.Name == "maven-central" {
			assert.Equal(t, "https://repo1.maven.org/maven2/", repo.Proxy.RemoteURL)
			assert.Equal(t, RepositoryTypeProxy, *repo.Type)
		}
	}
}

func TestMavenProxyRepositoryService_GetByName(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositoriesMockData).Build()
	defer s.teardown()
	repo, err := s.Client().MavenProxyRepositoryService.GetRepoByName("maven-central")
	assert.NoError(t, err)
	assert.Equal(t, "maven-central", repo.Name)
	assert.Equal(t, RepositoryTypeProxy, *repo.Type)
}

func TestMavenProxyRepositoryService_Add(t *testing.T) {
	s := newServerWrapper(t).Build()
	defer s.teardown()
	client := s.Client()
	apacheMavenRepoMockData.Name = apacheMavenRepoMockData.Name + "2"
	err := client.MavenProxyRepositoryService.Add(apacheMavenRepoMockData)
	assert.NoError(t, err)
	assertRemote(t, func() error {
		// on remote testing we can check if the repository was correctly inserted
		repos, err := client.MavenProxyRepositoryService.List()
		assert.NotEmpty(t, repos)
		for _, repo := range repos {
			if repo.Name == apacheMavenRepoMockData.Name {
				assert.Equal(t, "https://repo.maven.apache.org/maven2/", repo.Proxy.RemoteURL)
				return err
			}
		}
		assert.Fail(t, "Repository apache not found")
		return err
	})
}
