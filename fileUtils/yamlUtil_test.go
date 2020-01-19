package fileUtils

import (
	"os"
	"testing"
	"github.com/floge77/cloud2podcast/model"
)

func TestRadYamll(t *testing.T) {
	yamlUtil := YamlUtil{}

	config := yamlUtil.GetConfig("../config.yaml")

	if config.MinLength != 1800 {
		t.Errorf("MinLength should be 1800 but was %v", config.MinLength)
	}
	if config.Podcasts[0].Channel != "Q-Dance-Youtube" {
		t.Errorf("First Channel should be Q-Dance-Youtube but was %v", config.Podcasts[0].Channel)
	}
}

func TestWriteYaml(t *testing.T) {
	yamlUtil := YamlUtil{}
	podcast  := model.Podcast{
		Channel:               "myTestChannel",
		ChannelURL:            "mixcloud.com/mychannel",
		ChannelImageURL:       "myimage.com",
		PlaylistToDownloadURL: "my playlist",}

	podcasts := []*model.Podcast {
		&podcast,
	}
	config := &model.ConfigYaml {
		Podcasts: podcasts,
		MinLength: 123,
	} 
	
	yamlPath := "../test/myTestconfig.yaml"
	yamlUtil.WriteConfig(config, yamlPath)
	newConfig := yamlUtil.GetConfig(yamlPath)

	if newConfig.Podcasts[0].Channel != "myTestChannel" {
		t.Errorf("First Channel should be myTestChannel but was %v", config.Podcasts[0].Channel)
	}
	os.Remove(yamlPath)
}
