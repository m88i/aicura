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

type mvnGroupFakeService struct{}

var mvnGrpFake = make(map[string]*MavenGroupRepository)

func (m *mvnGroupFakeService) Update(repo MavenGroupRepository) error {
	if mvnGrpFake[repo.Name] == nil {
		return nil
	}
	mvnGrpFake[repo.Name] = &repo
	return nil
}

func (m *mvnGroupFakeService) List() ([]MavenGroupRepository, error) {
	repos := make([]MavenGroupRepository, 0)
	for _, repo := range mvnGrpFake {
		repos = append(repos, *repo)
	}
	return repos, nil
}

func (m *mvnGroupFakeService) GetRepoByName(name string) (*MavenGroupRepository, error) {
	return mvnGrpFake[name], nil
}
