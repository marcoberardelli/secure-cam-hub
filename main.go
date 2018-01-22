package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"
	"html/template"
)


func getIP() string {
	cmd := exec.Command("python3", "getip.py")
	ip, _ := cmd.CombinedOutput()
	return string(ip)
}

var templates *template.Template

func main() {
	//Load all the html files of the static/html folder
	templates = template.Must(template.ParseGlob("static/html/*.html"))



	r := mux.NewRouter()
	r.HandleFunc("/login", GetLogin).Methods("GET")
	r.HandleFunc("/login", PostLogin).Methods("POST")

	r.HandleFunc("/register", GetRegister).Methods("GET")
	r.HandleFunc("/register", PostRegister).Methods("POST")

	r.HandleFunc("/img/", GetImage).Methods("GET")
	r.HandleFunc("/img/", PostImage).Methods("POST")

	r.HandleFunc("/", Home).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
