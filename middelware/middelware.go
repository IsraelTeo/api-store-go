package middelware

import (
	"log"
	"net/http"
)

func Log(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Petición; %q, Método %q", r.URL.Path, r.Method)
		f(w, r)
	}
}
