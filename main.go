package main

import (
	"fmt"
	"github.com/floge77/cloud2podcastnew/reader"
)

func main() {
	yamlReader := YamlReader{}
	config := yamlReader.GetConfig("/Users/floge77/Development/github.com/c2pdownloader/config.yaml")
	fmt.Println(config)

}
