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

import "fmt"

// MavenGroupRepositoryService service to handle all Maven Proxy Repositories operations
type MavenGroupRepositoryService service

// Update updates a given Maven Group repository
func (m *MavenGroupRepositoryService) Update(repo MavenGroupRepository) error {
	req, err := m.client.put(m.client.appendVersion(fmt.Sprintf("/repositories/maven/group/%s", repo.Name)), "", repo)
	if err != nil {
		return err
	}
	_, err = m.client.do(req, nil)
	return err
}

// List lists all maven repositories from the Nexus Server
func (m *MavenGroupRepositoryService) List() ([]MavenGroupRepository, error) {
	repositories := []MavenGroupRepository{}
	err := m.client.mavenRepositoryService.list(RepositoryTypeGroup, &repositories)
	return repositories, err
}

// GetRepoByName gets the repository by name or nil if not found
func (m *MavenGroupRepositoryService) GetRepoByName(name string) (*MavenGroupRepository, error) {
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
