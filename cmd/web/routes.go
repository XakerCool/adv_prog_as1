package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/categories", app.byCategory)
	mux.HandleFunc("/edit", app.edit)
	mux.HandleFunc("/update", app.update)
	mux.HandleFunc("/delete", app.delete)
	mux.HandleFunc("/add", app.add)
	mux.HandleFunc("/add/create", app.create)

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}
