package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/lmagdanello/lmagdanello-blog/internal/loader"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var Templates *template.Template

func InitTemplates() {
	Templates = template.Must(template.New("").Funcs(template.FuncMap{
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
		"formatDatePT": func(t time.Time) string {
			mesesAbreviados := []string{
				"jan", "fev", "mar", "abr", "mai", "jun",
				"jul", "ago", "set", "out", "nov", "dez",
			}
			caser := cases.Title(language.BrazilianPortuguese)
			mes := caser.String(mesesAbreviados[int(t.Month())-1])
			return fmt.Sprintf("%02d %s %d", t.Day(), mes, t.Year())
		}}).ParseGlob("templates/*.html"))
}

func ExecuteTemplate(w io.Writer, name string, data any) error {
	return Templates.ExecuteTemplate(w, name, data)
}

func HomeHandler(posts []loader.Post, books []loader.Book, links []loader.SocialLink) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title":   "Home",
			"Content": posts,
			"Books":   books,
			"Links":   links,
		}
		err := Templates.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			log.Println("Erro no template:", err)
			http.Error(w, "Erro interno", 500)
		}
	}
}
