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

var allRepositories = `[ {
	"name" : "nuget-group",
	"format" : "nuget",
	"url" : "http://localhost:8081/repository/nuget-group",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : true
	},
	"group" : {
	  "memberNames" : [ "nuget-hosted", "nuget.org-proxy" ]
	},
	"type" : "group"
  }, {
	"name" : "maven-snapshots",
	"url" : "http://localhost:8081/repository/maven-snapshots",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : false,
	  "writePolicy" : "ALLOW"
	},
	"cleanup" : null,
	"maven" : {
	  "versionPolicy" : "SNAPSHOT",
	  "layoutPolicy" : "STRICT"
	},
	"format" : "maven2",
	"type" : "hosted"
  }, {
	"name" : "maven-central",
	"url" : "http://localhost:8081/repository/maven-central",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : false,
	  "writePolicy" : "ALLOW"
	},
	"cleanup" : null,
	"proxy" : {
	  "remoteUrl" : "https://repo1.maven.org/maven2/",
	  "contentMaxAge" : -1,
	  "metadataMaxAge" : 1440
	},
	"negativeCache" : {
	  "enabled" : true,
	  "timeToLive" : 1440
	},
	"httpClient" : {
	  "blocked" : false,
	  "autoBlock" : false,
	  "connection" : {
		"retries" : null,
		"userAgentSuffix" : null,
		"timeout" : null,
		"enableCircularRedirects" : false,
		"enableCookies" : false
	  },
	  "authentication" : null
	},
	"routingRuleName" : null,
	"maven" : {
	  "versionPolicy" : "RELEASE",
	  "layoutPolicy" : "PERMISSIVE"
	},
	"format" : "maven2",
	"type" : "proxy"
  }, {
	"name" : "nuget.org-proxy",
	"url" : "http://localhost:8081/repository/nuget.org-proxy",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : true,
	  "writePolicy" : "ALLOW"
	},
	"cleanup" : null,
	"proxy" : {
	  "remoteUrl" : "https://www.nuget.org/api/v2/",
	  "contentMaxAge" : 1440,
	  "metadataMaxAge" : 1440
	},
	"negativeCache" : {
	  "enabled" : true,
	  "timeToLive" : 1440
	},
	"httpClient" : {
	  "blocked" : false,
	  "autoBlock" : false,
	  "connection" : {
		"retries" : null,
		"userAgentSuffix" : null,
		"timeout" : null,
		"enableCircularRedirects" : false,
		"enableCookies" : false
	  },
	  "authentication" : null
	},
	"routingRuleName" : null,
	"nugetProxy" : {
	  "queryCacheItemMaxAge" : null
	},
	"format" : "nuget",
	"type" : "proxy"
  }, {
	"name" : "maven-releases",
	"url" : "http://localhost:8081/repository/maven-releases",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : false,
	  "writePolicy" : "ALLOW_ONCE"
	},
	"cleanup" : null,
	"maven" : {
	  "versionPolicy" : "RELEASE",
	  "layoutPolicy" : "STRICT"
	},
	"format" : "maven2",
	"type" : "hosted"
  }, {
	"name" : "nuget-hosted",
	"format" : "nuget",
	"url" : "http://localhost:8081/repository/nuget-hosted",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : true,
	  "writePolicy" : "ALLOW"
	},
	"cleanup" : null,
	"type" : "hosted"
  }, {
	"name" : "maven-public",
	"format" : "maven2",
	"url" : "http://localhost:8081/repository/maven-public",
	"online" : true,
	"storage" : {
	  "blobStoreName" : "default",
	  "strictContentTypeValidation" : true
	},
	"group" : {
	  "memberNames" : [ "maven-releases", "maven-snapshots", "maven-central" ]
	},
	"type" : "group"
  } ]`

func TestMavenProxyRepositoryService_List(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositories).Build()
	defer s.teardown()
	repos, err := s.Client().MavenProxyRepositoryService.List()
	assert.NoError(t, err)
	assert.Len(t, repos, 4)
	for _, repo := range repos {
		if repo.Name == "maven-central" {
			assert.Equal(t, "https://repo1.maven.org/maven2/", repo.Proxy.RemoteURL)
			assert.Equal(t, RepositoryTypeProxy, *repo.Type)
		}
	}
}

func TestMavenProxyRepositoryService_GetByName(t *testing.T) {
	s := newServerWrapper(t).WithResponse(allRepositories).Build()
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
	repository := MavenProxyRepository{
		Repository: Repository{
			Name:   "apache",
			Format: NewString(repositoryFormatMaven2),
			Type:   NewRepositoryType(RepositoryTypeProxy),
			Online: NewBool(true),
		},
		Proxy: Proxy{
			RemoteURL:      "https://repo.maven.apache.org/maven2/",
			ContentMaxAge:  1440,
			MetadataMaxAge: -1,
		},
		Storage: Storage{
			BlobStoreName:               "default",
			StrictContentTypeValidation: true,
		},
		NegativeCache: NegativeCache{
			Enabled:    true,
			TimeToLive: 1440,
		},
		HTTPClient: HTTPClient{
			AutoBlock: true,
			Blocked:   NewBool(false),
		},
		Maven: Maven{
			LayoutPolicy:  LayoutPolicyStrict,
			VersionPolicy: VersionPolicyRelease,
		},
	}
	err := client.MavenProxyRepositoryService.Add(repository)
	assert.NoError(t, err)
	assertRemote(t, func() error {
		// on remote testing we can check if the repository was correctly inserted
		repos, err := client.MavenProxyRepositoryService.List()
		assert.NotEmpty(t, repos)
		for _, repo := range repos {
			if repo.Name == "apache" {
				assert.Equal(t, "https://repo.maven.apache.org/maven2/", repo.Proxy.RemoteURL)
				return err
			}
		}
		assert.Fail(t, "Repository apache not found")
		return err
	})
}
