package tpl

import (
	"errors"
	"html/template"

	"github.com/plutov/games/pkg/config"
	"github.com/plutov/games/pkg/shared"

	"github.com/gocraft/web"
)

// Data struct
type Data struct {
	TemplateFile string
	LayoutFile   string
	Config       *config.Structure
	Data         interface{}
	Title        string
}

// Render func
func (s Data) Render(w web.ResponseWriter, r *web.Request) error {
	if s.TemplateFile == "" {
		tplErr := errors.New("tpl.Render requires non-empty TemplateFile")
		shared.LogErr(tplErr)
		return tplErr
	}

	if s.Title == "" {
		s.Title = "Games"
	}
	if s.LayoutFile == "" {
		s.LayoutFile = "templates/layouts/main.html"
	}

	t := template.Must(template.ParseFiles("templates/"+s.TemplateFile, s.LayoutFile))

	err := t.ExecuteTemplate(w, "base", s)
	shared.LogErr(err)

	return err
}
