#! /bin/bash

echo "Starting local cloud2podcast with dummy files"
pushd frontend && npm run build && popd
export downloadDir=test/ && export configYaml=config.yaml && go run main.go
