#! /bin/bash

echo "Starting local cloud2podcast with dummy files"
export downloadDir=test/ && export configYaml=config.yaml && go run main.go
