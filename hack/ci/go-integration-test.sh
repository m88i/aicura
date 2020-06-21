#!/bin/sh

# load image from cache to avoid pulling every time from internet
source ./hack/ci/load-image-cache.sh

# the base URL for the tests
export NEXUS_TEST_BASE_URL="http://localhost:8081/"

# spin up a container for the nexus server
docker run -d --rm -p 8081:8081 -e INSTALL4J_ADD_VM_PARAMS="-Dnexus.security.randompassword=false" docker.io/sonatype/nexus3:latest
# save in the file system for later runs
docker save -o ${IMAGE_OUTPUT} nexus3

# run the tests against the container
go test ./nexus/...
