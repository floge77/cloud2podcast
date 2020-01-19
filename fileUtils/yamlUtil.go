package fileUtils

import (
	"io/ioutil"
	"log"

	"github.com/floge77/cloud2podcast/model"
	"gopkg.in/yaml.v2"
)

type YamlUtil struct {
}

// GetConfig returns the config.yaml as config object
func (*YamlUtil) GetConfig(configYamlPath string) *model.ConfigYaml {
	config := model.ConfigYaml{}
	config = readYamlfile(configYamlPath, config)
	return &config
}

func (*YamlUtil) WriteConfig(config *model.ConfigYaml, configYamlPath string) {
	yamlFile, err := yaml.Marshal(config)
	if err != nil {
		log.Printf("Could not marshal %v", config)
		return
	}
	
	err =ioutil.WriteFile(configYamlPath, yamlFile, 0644)
	if err != nil {
		log.Printf("Could not write %v, to %v", config, configYamlPath)
	}
}

func readYamlfile(filePath string, config model.ConfigYaml) model.ConfigYaml {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("could not open %v error: %v", filePath, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config
}
