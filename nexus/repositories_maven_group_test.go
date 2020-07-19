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

func TestMavenGroupRepositoryService_Update(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositoriesMockData).Build()

	repos, err := s.Client().MavenGroupRepositoryService.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, repos)

	if len(repos) > 0 { //sanity check to not panic
		s = newServerWrapper(t).Build()
		apacheMavenRepoMockData.Name = apacheMavenRepoMockData.Name + "3"
		err := s.Client().MavenProxyRepositoryService.Add(apacheMavenRepoMockData)
		assert.NoError(t, err)

		repos[0].Group.MemberNames = append(repos[0].Group.MemberNames, apacheMavenRepoMockData.Name)
		err = s.Client().MavenGroupRepositoryService.Update(repos[0])
		assert.NoError(t, err)
	}
}

func TestMavenGroupRepositoryService_List(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositoriesMockData).Build()
	defer s.teardown()
	repos, err := s.Client().MavenGroupRepositoryService.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, repos)
	for _, repo := range repos {
		if repo.Name == "maven-public" {
			assert.Equal(t, RepositoryTypeGroup, *repo.Type)
		}
	}
}

func TestMavenGroupRepositoryService_GetByName(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositoriesMockData).Build()
	defer s.teardown()
	repo, err := s.Client().MavenGroupRepositoryService.GetRepoByName("maven-public")
	assert.NoError(t, err)
	assert.Equal(t, "maven-public", repo.Name)
	assert.Equal(t, RepositoryTypeGroup, *repo.Type)
}
