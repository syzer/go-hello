package lib

import (
	"net/http"
	"log"
	"path/filepath"
	"sync"
	"text/template"
)

func Serve() {

	http.Handle("/", &templateHandler{filename: "main.html"})

	if err := http.ListenAndServe(":3000", nil); err != nil {
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
	t.templ.Execute(w, nil)
}
