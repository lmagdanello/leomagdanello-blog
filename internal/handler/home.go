package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/lmagdanello/lmagdanello-blog/internal/loader"
)

var templates *template.Template

func InitTemplates() {
	templates = template.Must(template.New("").Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
		"truncate": func(length int, s string) string {
			if len(s) <= length {
				return s
			}
			return s[:length] + "..."
		},
		"now": func() time.Time {
			return time.Now()
		},
	}).ParseGlob("templates/*.html"))
}

func HomeHandler(posts []loader.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title":   "Home",
			"Content": posts,
		}

		err := templates.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			log.Println("Erro no template: base.html", err)
		}
	}
}
