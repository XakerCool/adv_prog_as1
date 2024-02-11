package main

import (
	"1_assignment/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	user *models.User
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
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, articles)
	if err != nil {
		app.serverError(w, err)
		return
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
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, articles)
	if err != nil {
		app.serverError(w, err)
		return
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
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, article)
	if err != nil {
		app.serverError(w, err)
		return
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
		app.serverError(w, err)
		return
	}
	fmt.Println(result)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	_, err := app.articles.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) add(w http.ResponseWriter, r *http.Request) {
	var files []string
	if user != nil {
		if user.Role == "student" {
			files = []string{
				"./ui/html/notAllowed.page.tmpl",
				"./ui/html/base.layout.tmpl",
			}
		} else if user.Role == "teacher" && !user.Approved {
			files = []string{
				"./ui/html/notAllowed.page.tmpl",
				"./ui/html/base.layout.tmpl",
			}
		}
	} else {
		files = []string{
			"./ui/html/login.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if user != nil {
		err = ts.Execute(w, user)
	} else {
		err = ts.Execute(w, nil)
	}

	if err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *application) create(w http.ResponseWriter, r *http.Request) {
	var createdArticle *models.Article
	err := json.NewDecoder(r.Body).Decode(&createdArticle)
	if err != nil {
		app.clientError(w, 400)
		return
	}

	result, err := app.articles.Insert(createdArticle.Category, createdArticle.Author, createdArticle.Readership, createdArticle.Title, createdArticle.Description, createdArticle.Content)
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Println(result)
}

func (app *application) getAllDeps(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/departments" {
		http.NotFound(w, r)
		return
	}
	articles, err := app.departments.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	files := []string{
		"./ui/html/departments.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, articles)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *application) createDep(w http.ResponseWriter, r *http.Request) {
	var createdArticle *models.Department
	err := json.NewDecoder(r.Body).Decode(&createdArticle)
	if err != nil {
		fmt.Println(err)
		app.clientError(w, 400)
		return
	}

	result, err := app.departments.Insert(createdArticle.DepName, createdArticle.StaffQuantity)
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Println(result)
}

func (app *application) loginPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/login.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var loginUser *models.User
	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		app.clientError(w, 400)
		return
	}
	result, err := app.users.Login(loginUser.Email, loginUser.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.clientError(w, 400)
			return
		} else {
			app.serverError(w, err)
			return
		}
	}
	if result != nil {
		user = result
	}
	http.Redirect(w, r, "/"+user.Role, 200)
}

func (app *application) registerPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/register.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var newUser *models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		app.clientError(w, 400)
		return
	}
	result, err := app.users.Register(newUser.FullName, newUser.Email, newUser.Role, newUser.Password, newUser.Approved)
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Println(result)
}

func (app *application) approve(w http.ResponseWriter, r *http.Request) {
	if user != nil {
		if user.Role == "admin" {
			files := []string{
				"./ui/html/approve.page.tmpl",
				"./ui/html/base.layout.tmpl",
			}
			ts, err := template.ParseFiles(files...)
			if err != nil {
				app.serverError(w, err)
				return
			}
			users, err := app.users.GetTeachers()
			if err != nil {
				app.serverError(w, err)
				return
			}
			err = ts.Execute(w, users)
			if err != nil {
				app.serverError(w, err)
				return
			}
		} else {
			http.Redirect(w, r, "/", 400)
			return
		}
	} else {
		http.Redirect(w, r, "/", 400)
		return
	}
}
