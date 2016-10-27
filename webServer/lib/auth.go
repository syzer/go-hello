package lib

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"strings"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		log.Println("no cookie... diet?")
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
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch provider {
	case "google":
		log.Println("Logging with google")
		break
	case "facebook":
		break
	case "github":
		break
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth provider %s not supported", provider)
	}

	switch action {
	case "login":
		log.Println("Handling login", provider)
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Providel not found", provider, err)
		}
		// Redirect, but don't modify state, and
		// dont provide additional parameter.
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("Error trying to get begining url", provider, err)
			w.WriteHeader(http.StatusTemporaryRedirect)
			w.Header().Set("Location", loginUrl)
		}

	case "callback":
		log.Println("Handling callback", provider)
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Error trying to get provider", provider, err)
		}
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("Error trying deserialize auth callback cookie", provider, err)
		}

		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("Error deserializing user from ", provider, err)
		}

		// TODO also use google email
		// if we store user data better to use sign cookies
		authCookie := objx.New(map[string]interface{}{
			"name":       user.Name(),
			"avatar_url": user.AvatarURL(),
		}).MustBase64()

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookie,
			Path:  "/",
		})

		w.WriteHeader(http.StatusTemporaryRedirect)
		w.Header()["Location"] = []string{"/chat"}

	default:
		fmt.Fprintf(w, "Auth action %s not supported", action)
		w.WriteHeader(http.StatusNotFound)
	}
}
