package web

import (
	"encoding/json"
	"net/http"

	"github.com/plutov/games/pkg/env"
	"github.com/plutov/games/pkg/shared"

	"github.com/gocraft/web"
	"github.com/gorilla/context"
)

// Context struct
type Context struct {
	Env env.Context
}

// ListenAndServe func
func ListenAndServe(e env.Context) {
	ctx := new(Context)
	ctx.Env = e

	r := web.New(Context{}).
		Get("/", ctx.home).
		Get("/game", ctx.game)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)
	staticFolders := []string{"scripts", "styles"}
	for _, sf := range staticFolders {
		serveMux.Handle("/"+sf+"/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	}

	shared.Infof("Starting on http://%s", ctx.Env.Config.Addr)
	err := http.ListenAndServe(ctx.Env.Config.Addr, context.ClearHandler(serveMux))

	shared.LogErr(err)
}

// Ajax func
func (c *Context) Ajax(w web.ResponseWriter, r *web.Request, response interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resultJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJSON)
}
