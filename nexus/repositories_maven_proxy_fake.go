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

type mvnProxyFakeService struct{}

var mvnProxyFake = make(map[string]*MavenProxyRepository)

func (m *mvnProxyFakeService) Add(repositories ...MavenProxyRepository) error {
	for _, repo := range repositories {
		mvnProxyFake[repo.Name] = &repo
	}
	return nil
}

func (m *mvnProxyFakeService) List() ([]MavenProxyRepository, error) {
	repos := make([]MavenProxyRepository, 0)
	for _, repo := range mvnProxyFake {
		repos = append(repos, *repo)
	}
	return repos, nil
}

func (m *mvnProxyFakeService) GetRepoByName(name string) (*MavenProxyRepository, error) {
	return mvnProxyFake[name], nil
}
