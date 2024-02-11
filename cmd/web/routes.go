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

	mux.HandleFunc("/departments", app.getAllDeps)
	mux.HandleFunc("/departments/add", app.createDep)

	mux.HandleFunc("/login_page", app.loginPage)
	mux.HandleFunc("/login_page/login", app.login)
	mux.HandleFunc("/register_page", app.registerPage)
	mux.HandleFunc("/register_page/register", app.register)

	mux.HandleFunc("/admin/approve", app.approve)

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}
