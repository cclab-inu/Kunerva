#!/bin/bash

# remove old images
docker images | grep ubuntu-autopol | awk '{print $3}' | xargs -I {} docker rmi -f {} 2> /dev/null

# login docker hub
docker login

if [ $? -ne 0 ]; then
    exit
fi

# create new images
docker build --tag cclabinu/ubuntu-kunerva:latest .

# push new images
docker push cclabinu/ubuntu-kunerva:latest
