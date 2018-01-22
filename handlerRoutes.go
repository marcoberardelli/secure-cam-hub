package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func GetRegister(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "register.html", nil)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {


	templates.ExecuteTemplate(w, "index.html", nil)
}

func PostRegister(w http.ResponseWriter, r *http.Request) {


	templates.ExecuteTemplate(w, "index.html", nil)
}