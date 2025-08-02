package httpserver

import (
	"fmt"
	"net/http"

	"github.com/quii/go-specs-greet/domain/interactions"
)

var (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, Reply(interactions.Greet))
	mux.HandleFunc(cursePath, Reply(interactions.Curse))
	return mux
}

func Reply(f func(name string) (interaction string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
