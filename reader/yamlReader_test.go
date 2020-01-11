package reader

import "testing"

func TestYamlReader(t *testing.T) {
	reader := YamlReader{}
	
	config := reader.GetConfig("../config.yaml")

	if config.MinLength != 1800 {
		t.Errorf("MinLength should be 1800 but was %v", config.MinLength)
	}
	if config.Podcasts[0].Channel != "Q-Dance-Youtube" {
		t.Errorf("First Channel should be Q-Dance-Youtube but was %v", config.Podcasts[0].Channel)
	}
}
