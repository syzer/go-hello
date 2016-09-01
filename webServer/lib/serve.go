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
)

func Serve() {
	var port = flag.String("addr", ":3000", "Address of the application")
	flag.Parse()

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/room", r)

	http.Handle("/assets", http.StripPrefix("/assets",
		http.FileServer(http.Dir("assets"))))

	http.Handle("/login", &templateHandler{filename: "login.html"})

	// get the room, cooroutine/thread
	go r.run()
	log.Println("Starting on: %s", *port)
	// start on 3000
	if err := http.ListenAndServe(*port, nil); err != nil {
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
