package web

import (
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
		Middleware(web.LoggerMiddleware).
		Get("/", ctx.home)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", r)
	staticFolders := []string{"scripts"}
	for _, sf := range staticFolders {
		serveMux.Handle("/"+sf+"/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	}

	shared.Infof("Starting on http://%s", ctx.Env.Config.Addr)
	err := http.ListenAndServe(ctx.Env.Config.Addr, context.ClearHandler(serveMux))

	shared.LogErr(err)
}
