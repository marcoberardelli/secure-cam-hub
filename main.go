package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os/exec"
	"html/template"
	"github.com/marcoberardelli/secure-cam/handler"
)


var templates *template.Template


func handlerHome(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}


func getIP() string {
	cmd := exec.Command("python3", "getip.py")
	ip, _ := cmd.CombinedOutput()
	return string(ip)
}


func main() {
	//Load all the html files in the static folder
	templates = template.Must(template.ParseGlob("static/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/img/", handler.GetImage).Methods("GET")
	r.HandleFunc("/img/", handler.PostImage).Methods("POST")
	r.HandleFunc("/", handlerHome).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
