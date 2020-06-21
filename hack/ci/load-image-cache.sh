#!/bin/sh
set -eu

IMAGE_OUTPUT_FILENAME=${IMAGE_OUTPUT}/nexus3-latest.tar
if [ -f ${IMAGE_OUTPUT_FILENAME} ]; then
    echo "Found image file ${IMAGE_OUTPUT_FILENAME}, loading into internal Docker Registry"
    docker load -i ${IMAGE_OUTPUT_FILENAME}
else
    echo "Image not found on ${IMAGE_OUTPUT_FILENAME}, it needs to be pulled from an external registry"
fi
