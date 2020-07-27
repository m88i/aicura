![Nexus Client Integration Checks](https://github.com/m88i/aicura/workflows/Nexus%20Client%20Integration%20Checks/badge.svg)
[![codecov](https://codecov.io/gh/m88i/aicura/branch/master/graph/badge.svg)](https://codecov.io/gh/m88i/aicura)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/m88i/nexus-operator)
[![License: LGPL v3](https://img.shields.io/badge/License-LGPL%20v3-blue.svg)](https://www.gnu.org/licenses/lgpl-3.0)
 
# AICURA: A Sonatype Nexus API in Go

`aicura` is a client library for [Sonatype Nexus Manager v3 API](https://help.sonatype.com/repomanager3/rest-and-integration-api) written in GoLang.

## How to Use

The most straightforward way to use this library is creating a new [`Client`](https://github.com/m88i/aicura/blob/master/nexus/client.go) and make calls to the API through services:

```go
import   "github.com/m88i/aicura/nexus"

(...)

// New client with default credentials
client := nexus.NewClient("http://localhost:8081").WithCredentials("admin", "admin123").Build()
user, err := client.UserService.GetUserByID("admin")
if err != nil {
    return err
}
print(user.Name)
```

### Fake Client

To use this library in your unit tests, create an instance of a "fake" `Client` instead:

```go
import   "github.com/m88i/aicura/nexus"

(...)

client := nexus.NewFakeClient()

// all interfaces remain the same
user, err := client.UserService.GetUserByID("admin")
if err != nil {
    return err
}
print(user.Name) //will print nothing since there's no user in the cache, call client.UserService.Add(user) first :)
```

The fake `Client` is backed up by hash maps, so all write and read operations will work like in real scenarios.

## Development

To run a local Nexus Server 3.x container with Podman:

```shell
$ podman run --rm -it -p 8081:8081 -e INSTALL4J_ADD_VM_PARAMS="-Dnexus.security.randompassword=false" docker.io/sonatype/nexus3
```

Use Postman or even the web browser to play around with the Nexus REST API at http://localhost:8081 endpoint. The default credentials are `admin`/`admin123`.