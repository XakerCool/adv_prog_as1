package main

import (
	"1_assignment/pkg/models"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	articles, err := app.articles.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, articles)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) byCategory(w http.ResponseWriter, r *http.Request) {
	readership := r.URL.Query().Get("readership")
	if readership != "students" && readership != "staff" && readership != "applicants" {
		http.NotFound(w, r)
		return
	}
	articles, err := app.articles.GetByCategory(readership)
	if err != nil {
		app.serverError(w, err)
		return
	}
	files := []string{
		"./ui/html/for_" + readership + ".page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, articles)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
func (app *application) edit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	article, err := app.articles.GetById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	files := []string{
		"./ui/html/edit.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, article)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
func (app *application) update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var updatedArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&updatedArticle)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	result, err := app.articles.Update(id, updatedArticle.Category, updatedArticle.Author, updatedArticle.Title, updatedArticle.Description, updatedArticle.Content, updatedArticle.Readership)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	fmt.Println(result)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	_, err := app.articles.Delete(id)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}

}
func (app *application) add(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
func (app *application) create(w http.ResponseWriter, r *http.Request) {
	var createdArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&createdArticle)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
		return
	}

	result, err := app.articles.Insert(createdArticle.Category, createdArticle.Author, createdArticle.Readership, createdArticle.Title, createdArticle.Description, createdArticle.Content)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	fmt.Println(result)
}
