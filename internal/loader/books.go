package loader

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Book struct {
	Title  string `yaml:"title"`
	Author string `yaml:"author"`
	Link   string `yaml:"link"`
}

type BookData struct {
	Books []Book `yaml:"books"`
}

func LoadBooks(filePath string) ([]Book, error) {
	var data BookData
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}
	return data.Books, nil
}
