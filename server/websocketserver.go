package server

import (
	"log"
	"net/http"
	"text/template"

	"github.com/pdxjohnny/easysock"
)

var homeTempl = template.Must(template.ParseFiles("../static/home.html"))

func ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func Run() error {
	go easysock.Hub.Run()
	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/ws", easysock.ServeWs)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
