package lib

import (
	"net/http"
	"strings"
	"log"
	"fmt"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		// no cookie... diet?
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		// i give up and panic
		panic(err.Error())
	} else {
		// next() - in express.js
		h.next.ServeHTTP(w, r)
	}
}

// AKA require auth
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

// /auth/{action=login|callback}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]

	switch provider {
	case "google":
		break;
	case "facebook":
		break;
	case "github":
		break;
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth provider %s not supported", provider)
	}

	switch action {
	case "login":
		log.Println("Handle login", provider)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
