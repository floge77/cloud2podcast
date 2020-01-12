#! /bin/bash

IP=`ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1'`
echo $IP
docker run --rm --name cloud2podcast -p 8080:8080 -v $HOME/Desktop/test:/downloads -e HOST_IP=localhost -it floge77/cloud2podcast
