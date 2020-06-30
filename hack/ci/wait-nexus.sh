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

set -eu

# See: https://gist.github.com/rgl/f90ff293d56dbb0a1e0f7e7e89a81f42

declare -r PROBE=${NEXUS_TEST_BASE_URL}/service/rest/v1/status

wait-for-url() {
    echo "Waiting for $1"
    timeout -s TERM 5m bash -c \
        'while [[ "$(curl -s -o /dev/null -L -w ''%{http_code}'' ${0})" != "200" ]];\
    do echo "Waiting for ${0}" && sleep 5;\
    done' ${1}
    echo "OK!"
    curl -I $1
}
wait-for-url ${PROBE}
