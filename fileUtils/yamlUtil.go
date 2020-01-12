package fileUtils

import (
	"io/ioutil"
	"log"

	"github.com/floge77/cloud2podcast/model"
	"gopkg.in/yaml.v2"
)

type YamlReader struct {
}

func (*YamlReader) GetConfig(configYamlPath string) *model.PodcastConfigYaml {
	config := model.PodcastConfigYaml{}
	config = readYamlfile(configYamlPath, config)
	return &config
}

func readYamlfile(filePath string, config model.PodcastConfigYaml) model.PodcastConfigYaml {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Could not open %v Error: %v", filePath, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
