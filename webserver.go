package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"mime"
	"net/http"
	"net/url"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func startWebServer() {
	// Workaround for Windows
	mime.AddExtensionType(".js", "text/javascript")

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/scoretable", getScoretableHandler)

	log.Printf("Starting web server at http://localhost:%d/", cfg.Port)
	http.ListenAndServe(":"+fmt.Sprintf("%d", cfg.Port), mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func getScoretableHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := u.Query()
	role := params.Get("role")

	var res *Scoretable

	switch {
	case role == "tank":
		res = table.Tank
	case role == "dps":
		res = table.Dps
	case role == "sup":
		res = table.Support
	default:
		http.Error(w, "Invalid role", http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(res)
	w.Write(data)
}
