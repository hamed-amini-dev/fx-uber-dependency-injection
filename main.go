package main

import (
	"context"
	"net/http"

	"go.uber.org/fx"
)

// fx.Invoke invoke a function after all of the dependencies for an application have been provided
func main() {
	fx.New(
		fx.Provide(http.NewServeMux),
		fx.Invoke(New),
		fx.Invoke(register),
	).Run()
}

// fx.Lifecycle is a package that provides a way to manage the lifecycle of components
func register(
	lifecycle fx.Lifecycle, mux *http.ServeMux,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go http.ListenAndServe(":8080", mux)
				return nil
			},
		},
	)
}

// App for http requests
type App struct {
	mux *http.ServeMux
}

// New http handler
func New(s *http.ServeMux) *App {
	h := App{s}
	h.routes()
	return &h
}

// Register Routes for all http endpoints
func (h *App) routes() {
	h.mux.HandleFunc("/", h.HelloWorld)
}

func (h *App) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}
