package handler

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/lmagdanello/lmagdanello-blog/internal/loader"
)

func PostHandler(posts []loader.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := filepath.Base(r.URL.Path)

		for _, p := range posts {
			if p.Filename == slug {
				err := templates.ExecuteTemplate(w, "base.html", map[string]interface{}{
					"Content": p,       // Dados do post
					"Title":   p.Title, // Título da página
				})
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
