package middelware

import (
	"log"
	"net/http"

	"github.com/IsraelTeo/api-store-go/authorization"
)

func Log(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Petición; %q, Método %q", r.URL.Path, r.Method)
		f(w, r)
	}
}

func Authentication(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := authorization.ValidateToken(token)
		if err != nil {
			forbidden(w, r)
			return
		}

		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application(json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("No tiene autorización"))
}
