#! /bin/bash

HOST_IP=`ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1'`

docker run --rm --name cloud2podcast -p 80:80 -v $HOME/downloads:/downloads -e downloadDir=$HOME/downloads -e HOST_IP=$HOST_IP -it cloud2podcast
