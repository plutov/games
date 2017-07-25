package web

import (
	"github.com/plutov/games/pkg/tpl"

	"fmt"
	"time"

	"github.com/gocraft/web"
)

func (c *Context) home(w web.ResponseWriter, r *web.Request) {
	sessionID := time.Now().UnixNano()
	sessionIDStr := fmt.Sprint(sessionID)

	c.Env.GM.InitSession(sessionIDStr)

	d := tpl.Data{
		Config:       c.Env.Config,
		TemplateFile: "pages/home.html",
		Data: struct {
			SessionID string
		}{
			SessionID: sessionIDStr,
		},
	}

	d.Render(w, r)
}

func (c *Context) game(w web.ResponseWriter, r *web.Request) {
	userInput := r.FormValue("input")
	sessionID := r.FormValue("session_id")

	answer, err := c.Env.GM.Reply(sessionID, userInput)
	if err != nil {
		answer = err.Error()
	}

	c.Ajax(w, r, struct {
		Answer string `json:"answer"`
	}{
		Answer: answer,
	})
}
