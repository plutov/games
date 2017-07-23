package web

import (
	"github.com/plutov/games/pkg/game"
	"github.com/plutov/games/pkg/shared"
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

func (c *Context) game(w web.ResponseWriter, r *web.Request) {
	userInput := r.FormValue("input")
	answer, err := game.Reply(userInput)
	shared.LogErr(err)
	c.Ajax(w, r, answer)
}
