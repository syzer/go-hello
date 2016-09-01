package lib

import "net/http"

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
