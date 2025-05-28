package main

import (
	"log"
	"net/http"

	"github.com/lmagdanello/lmagdanello-blog/internal/handler"
	"github.com/lmagdanello/lmagdanello-blog/internal/loader"
)

func main() {
	// Inicializa os templates
	handler.InitTemplates()

	// Carregar os livros currently reading
	books, err := loader.LoadBooks("data/books.yaml")
	if err != nil {
		log.Fatalf("Erro ao carregar livros: %v", err)
	}

	// Social Media links
	links, err := loader.LoadLinks("data/links.yaml")
	if err != nil {
		log.Fatalf("Erro ao carregar links: %v", err)
	}

	// Carregar os posts em Markdown
	posts, err := loader.LoadPosts("posts")
	if err != nil {
		log.Fatalf("Erro ao carregar posts: %v", err)
	}

	// Rotas
	http.HandleFunc("/", handler.HomeHandler(posts, books, links))
	http.HandleFunc("/post/", handler.PostHandler(posts, books, links))

	log.Println("Servidor rodando em http://localhost:8080")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
