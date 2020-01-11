package reader

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/floge77/cloud2podcastnew/model"
	"gopkg.in/yaml.v2"
)

type YamlReader struct {
}

func (*YamlReader) GetConfig(yamlPath string) PodcastConfigYaml {
	config := PodcastConfigYaml{}
	config = ReadYamlfile(yamlPath, config)
	return config
}

func readYamlfile(filePath string, config PodcastConfigYaml) PodcastConfigYaml {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Could not open %v Error: %v", filePath, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", config)
	return config
}
