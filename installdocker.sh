#!/usr/bin/env bash
wget https://download.docker.com/linux/static/stable/x86_64/docker-18.03.1-ce.tgz

mkdir -p ./gen/docker

tar xzvf ./docker-18.03.1-ce.tgz 

cp ./docker/* /usr/bin/

dockerd &

