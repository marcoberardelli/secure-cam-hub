package handlerRoutes

import (
	"net/http"
	"html/template"
)

type HandlerRoutes struct {
	Templates *template.Template
}

func (handler *HandlerRoutes) Home(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "index.html", nil)
}

func (handler *HandlerRoutes) GetLogin(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "login.html", nil)
}

func (handler *HandlerRoutes) GetRegister(w http.ResponseWriter, r *http.Request) {
	handler.Templates.ExecuteTemplate(w, "register.html", nil)
}

func (handler *HandlerRoutes) PostLogin(w http.ResponseWriter, r *http.Request) {


	handler.Templates.ExecuteTemplate(w, "index.html", nil)
}

func (handler *HandlerRoutes) PostRegister(w http.ResponseWriter, r *http.Request) {


	handler.Templates.ExecuteTemplate(w, "index.html", nil)
}