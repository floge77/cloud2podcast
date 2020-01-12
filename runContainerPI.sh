#! /bin/bash

HOST_IP=`ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1'`

docker run --rm --name cloud2podcast -p 80:8080 -v $HOME/downloads:/downloads -e HOST_IP="cloud2podcast" -e port=80 -it floge77/cloud2podcastpi
