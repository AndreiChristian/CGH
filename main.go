package main

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	c.value--
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {

	counter := &Counter{}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	r.Post("/increase", func(w http.ResponseWriter, _ *http.Request) {
		counter.Increase()
		tmplString := "<div class=\"text-zinc-900 text-xl\" id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplString))
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)
	})

	r.Post("/decrease", func(w http.ResponseWriter, _ *http.Request) {
		counter.Increase()
		tmplString := "<div class=\"text-zinc-900 text-xl\" id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplString))
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}
		tmpl.ExecuteTemplate(w, "counter", data)
	})

	http.ListenAndServe(":4321", r)
}
