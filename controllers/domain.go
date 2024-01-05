package controllers

import (
	"html/template"
	"net/http"
	"ulfr/models"

	"github.com/julienschmidt/httprouter"
)

type Domain struct {
	// ...
}

func (domain Domain) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("domain/list")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Domains"], err = dbi.FindAll([]models.Domain{})
	if err != nil {
		panic(err)
	}

	view.ExecuteTemplate(w, "template", data)
}

func (domain Domain) Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("domain/add")...)
	if err != nil {
		panic(err)
	}

	view.ExecuteTemplate(w, "template", nil)
}

func (domain Domain) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	name := r.FormValue("name")
	description := r.FormValue("description")

	d := &models.Domain{
		Name:        name,
		Description: description,
	}

	err := dbi.Create(d)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/domain", http.StatusSeeOther)
}

func (domain Domain) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	d, err := dbi.Find([]models.Domain{}, map[string]interface{}{"id": ps.ByName("id")})
	if err != nil {
		panic(err)
	}

	dbi.Delete(d)

	http.Redirect(w, r, "/domain", http.StatusSeeOther)
}

func (domain Domain) Update(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	id := Params.ByName("id")
	name := r.FormValue("name")
	description := r.FormValue("description")

	d, err := dbi.Find([]models.Domain{}, map[string]interface{}{"id": id})

	if err != nil {
		panic(err)
	}

	dbi.Updates(d, models.Domain{
		Name:        name,
		Description: description,
	})

	http.Redirect(w, r, "/domain/update/"+id, http.StatusSeeOther)
}

func (domain Domain) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("domain/get")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Domain"], err = dbi.Find([]models.Domain{}, map[string]interface{}{"id": ps.ByName("id")})

	if err != nil {
		panic(err)
	}
	view.ExecuteTemplate(w, "template", data)
}
