#!/bin/bash
set -eu

# See: https://gist.github.com/rgl/f90ff293d56dbb0a1e0f7e7e89a81f42

declare -r PROBE=${NEXUS_TEST_BASE_URL}/service/rest/v1/status/check

wait-for-url() {
    echo "Waiting for $1"
    timeout -s TERM 5m bash -c \
        'while [[ "$(curl -s -o /dev/null -L -w ''%{http_code}'' ${0})" != "200" ]];\
    do echo "Waiting for ${0}" && sleep 5;\
    done' ${1}
    echo "OK!"
    curl -I $1
}
wait-for-url http://${PROBE}
