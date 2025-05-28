package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/lmagdanello/lmagdanello-blog/internal/loader"
)

func PostHandler(posts []loader.Post, books []loader.Book, links []loader.SocialLink) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := filepath.Base(r.URL.Path)

		for _, p := range posts {
			if p.Filename == slug {
				data := map[string]interface{}{
					"Title":   p.Title,
					"Content": p,
					"Books":   books,
					"Links":   links,
				}
				err := Templates.ExecuteTemplate(w, "base.html", data)
				if err != nil {
					log.Printf("Erro ao renderizar post: %v", err)
					http.Error(w, "Erro interno", http.StatusInternalServerError)
				}
				return
			}
		}
		http.NotFound(w, r)
	}
}
