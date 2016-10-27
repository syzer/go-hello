package lib

import (
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"github.com/syzer/go-hello/webServer/lib/trace"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

func Serve() {
	var addr = flag.String("addr", ":3000", "Address of the application")
	//TODO use protocol
	//var protocol = flag.String("protocol", "http", "Address of the application")
	var googleSecret = flag.String("secret", "no-way", "Please provide google secret")
	var googleClientId = flag.String("client", "no-way", "Please provide google clientId")
	flag.Parse()

	//TODO extract cont
	gomniauth.SetSecurityKey("14DOURGWzy2ZkagebOHXC9TS7PEZ6j")
	gomniauth.WithProviders(
		google.New(*googleClientId, *googleSecret, "http://127.0.0.1:8000/auth/callback/google"),
		github.New("key", "secret", "http://auth/callback/github"),
		facebook.New("key", "secret", "http://auth/callback/facebook"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth", loginHandler)

	http.Handle("/room", r)

	// get the room, cooroutine/thread
	go r.run()
	log.Printf("Starting on http://localhost%v", *addr)
	// start on 3000
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		// if executable then relative part gives problems!!!
		t.templ = template.Must(template.ParseFiles(filepath.Join(".", "html", t.filename)))
	})

	data := map[string]interface{}{
		"Host": r.Host,
	}
	// added data in runtime
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	// pass request as data object
	t.templ.Execute(w, data)
}
