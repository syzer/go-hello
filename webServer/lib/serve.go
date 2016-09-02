package lib

import (
	"net/http"
	"log"
	"path/filepath"
	"sync"
	"text/template"
	"flag"
	"os"
	"github.com/syzer/go-hello/webServer/lib/trace"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/gomniauth/providers/github"
)

func Serve() {
	var addr = flag.String("addr", ":3000", "Address of the application")
	var protocol = flag.String("protocol", "http", "Address of the application")
	flag.Parse()

	//TODO extract cont
	gomniauth.SetSecurityKey("14DOURGWzy2ZkagebOHXC9TS7PEZ6j")
	gomniauth.WithProviders(
		google.New("key", "secret", protocol + "://" + addr + "/auth/callback/google"),
		github.New("key", "secret", protocol + "://" + addr + "/auth/callback/github"),
		facebook.New("key", "secret", protocol + "://" + addr + "/auth/callback/facebook"),
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
	log.Println("Starting on: %s", *addr)
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
		// if execucutable then relative part gives problems!!!
		t.templ = template.Must(template.ParseFiles(filepath.Join(".", "html", t.filename)))
	})
	// pass request as data object
	t.templ.Execute(w, r)
}
