#!/bin/bash
#     Copyright 2020 Aicura Nexus Client and/or its authors
#
#     This file is part of Aicura Nexus Client.
#
#     Aicura Nexus Client is free software: you can redistribute it and/or modify
#     it under the terms of the GNU Lesser General Public License as published by
#     the Free Software Foundation, either version 3 of the License, or
#     (at your option) any later version.
#
#     Aicura Nexus Client is distributed in the hope that it will be useful,
#     but WITHOUT ANY WARRANTY; without even the implied warranty of
#     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#     GNU Lesser General Public License for more details.
#
#     You should have received a copy of the GNU Lesser General Public License
#     along with Aicura Nexus Client.  If not, see <https://www.gnu.org/licenses/>.

set -eux

# load image from cache to avoid pulling every time from internet
source ./hack/ci/load-image-cache.sh

# the base URL for the tests
export NEXUS_TEST_BASE_URL="http://localhost:8081"

# spin up a container for the nexus server
docker run -d --rm -p 8081:8081 -e INSTALL4J_ADD_VM_PARAMS="-Dnexus.security.randompassword=false" docker.io/sonatype/nexus3:latest
# save in the file system for later runs
mkdir -p ${IMAGE_OUTPUT}
docker save -o ${IMAGE_OUTPUT_FILENAME} nexus3

source ./hack/ci/wait-nexus.sh

# run the tests against the container
go test ./nexus/...
