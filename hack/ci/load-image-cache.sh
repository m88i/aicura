#!/bin/sh

if [ -f ${IMAGE_OUTPUT} ]; then
    echo "Found image file ${IMAGE_OUTPUT}, loading into internal Docker Registry"
    docker load -i ${IMAGE_OUTPUT}
else
    echo "Image not found on ${IMAGE_OUTPUT}, it needs to be pulled from an external registry"
fi
