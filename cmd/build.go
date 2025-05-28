package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lmagdanello/lmagdanello-blog/internal/handler"
	"github.com/lmagdanello/lmagdanello-blog/internal/loader"
)

func buildStaticSite() {
	// Inicializa os templates
	handler.InitTemplates()

	// Carrega dados
	books, _ := loader.LoadBooks("data/books.yaml")
	links, _ := loader.LoadLinks("data/links.yaml")
	posts, _ := loader.LoadPosts("posts")

	// Cria diretório de saída
	os.MkdirAll("public/post", os.ModePerm)

	// Gera página inicial
	homeFile, err := os.Create("public/index.html")
	if err != nil {
		log.Fatalf("Erro ao criar index.html: %v", err)
	}
	defer homeFile.Close()

	err = handler.ExecuteTemplate(homeFile, "base.html", map[string]interface{}{
		"Title":   "Home",
		"Content": posts,
		"Books":   books,
		"Links":   links,
	})
	if err != nil {
		log.Fatalf("Erro ao renderizar home: %v", err)
	}

	// Gera páginas de post individual
	for _, post := range posts {
		f, err := os.Create(filepath.Join("public/post", post.Filename+".html"))
		if err != nil {
			log.Printf("Erro ao criar arquivo para %s: %v", post.Filename, err)
			continue
		}
		err = handler.ExecuteTemplate(f, "base.html", map[string]interface{}{
			"Title":   post.Title,
			"Content": post,
			"Books":   books,
			"Links":   links,
		})
		if err != nil {
			log.Printf("Erro ao renderizar post %s: %v", post.Filename, err)
		}
		f.Close()
	}
	fmt.Println("Site estático gerado com sucesso em ./public/")
}
