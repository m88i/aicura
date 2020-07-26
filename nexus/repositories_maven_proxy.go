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

	"golang.org/x/sync/errgroup"
)

// MavenProxyRepositoryService service to handle all Maven Proxy Repositories operations
type MavenProxyRepositoryService interface {
	Add(repositories ...MavenProxyRepository) error
	List() ([]MavenProxyRepository, error)
	GetRepoByName(name string) (*MavenProxyRepository, error)
}

type mavenProxyRepositoryService service

// Add adds new Proxy Maven repositories to the Nexus Server
func (m *mavenProxyRepositoryService) Add(repositories ...MavenProxyRepository) error {
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
func (m *mavenProxyRepositoryService) List() ([]MavenProxyRepository, error) {
	repositories := []MavenProxyRepository{}
	err := m.client.mavenRepositoryService.list(RepositoryTypeProxy, &repositories)
	return repositories, err
}

// GetRepoByName gets the repository by name or nil if not found
func (m *mavenProxyRepositoryService) GetRepoByName(name string) (*MavenProxyRepository, error) {
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
