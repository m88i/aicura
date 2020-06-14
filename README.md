# AICURA: A Sonatype Nexus API in Go

`aicura` is a client library for [Sonatype Nexus Manager v3 API](https://help.sonatype.com/repomanager3/rest-and-integration-api) written in GoLang.

## Development

To run a local Nexus Server 3.x container with Podman:

```shell
$ podman run --rm -it -p 8081:8081 -e INSTALL4J_ADD_VM_PARAMS="-Dnexus.security.randompassword=false" docker.io/sonatype/nexus3
```

Use Postman or even the web browser to play around with the Nexus REST API at http://localhost:8081 endpoint. The default credentials are `admin`/`admin123`.