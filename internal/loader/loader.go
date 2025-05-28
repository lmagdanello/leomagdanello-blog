package loader

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v2"
)

type Post struct {
	Title    string
	Author   string
	Date     time.Time
	Tags     string
	Filename string
	HTML     string
}

type frontMatter struct {
	Title  string   `yaml:"title"`
	Author string   `yaml:"author"`
	Date   string   `yaml:"date"`
	Tags   []string `yaml:"tags"`
}

// LoadPosts carrega os posts em Markdown
// do diretorio e retorna uma slice de Post
func LoadPosts(dir string) ([]Post, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("Erro ao ler o diretorio de posts!")
		return nil, err
	}

	var posts []Post
	md := goldmark.New()

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		content, err := os.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		parts := strings.SplitN(string(content), "---", 3)
		if len(parts) < 3 {
			log.Printf("arquivo %s inválido: sem frontmatter", file.Name())
			continue
		}

		var meta frontMatter
		if err := yaml.Unmarshal([]byte(parts[1]), &meta); err != nil {
			log.Printf("Erro ao parsear YAML de %s: %v", file.Name(), err)
			continue
		}

		parsedDate, err := time.Parse("2006-01-02", meta.Date)
		if err != nil {
			log.Printf("Data inválida em %s: %v", file.Name(), err)
			continue
		}

		var html strings.Builder
		if err := md.Convert([]byte(parts[2]), &html); err != nil {
			log.Printf("Erro ao converter markdown de %s: %v", file.Name(), err)
			continue
		}

		posts = append(posts, Post{
			Title:    meta.Title,
			Author:   meta.Author,
			Date:     parsedDate,
			Tags:     strings.Join(meta.Tags, ", "),
			Filename: strings.TrimSuffix(file.Name(), ".md"),
			HTML:     html.String(),
		})
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts, nil
}
