package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Dashboard struct {
	// ...
}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("dashboard")...)
	if err != nil {
		panic(err)
	}
	view.ExecuteTemplate(w, "template", nil)
}
