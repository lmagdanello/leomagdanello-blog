package loader

import (
	"os"

	"gopkg.in/yaml.v2"
)

type SocialLink struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
	Icon string `yaml:"icon"`
	Alt  string `yaml:"alt"`
}

type LinkData struct {
	Links []SocialLink `yaml:"links"`
}

func LoadLinks(filePath string) ([]SocialLink, error) {
	var data LinkData
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data.Links, nil
}
