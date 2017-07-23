package web

import (
	"github.com/plutov/games/pkg/tpl"

	"github.com/gocraft/web"
)

func (c *Context) home(w web.ResponseWriter, r *web.Request) {
	d := tpl.Data{
		Config:       c.Env.Config,
		TemplateFile: "pages/home.html",
	}

	d.Render(w, r)
}
