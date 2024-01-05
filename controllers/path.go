package controllers

import (
	"html/template"
	"net/http"
	"strconv"
	"ulfr/models"

	"github.com/julienschmidt/httprouter"
)

type Path struct {
	// ...
}

func (path Path) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("path/list")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Paths"], err = dbi.FindAll([]models.Path{})
	if err != nil {
		panic(err)
	}
	data["Domains"], err = dbi.FindAll([]models.Domain{})
	if err != nil {
		panic(err)
	}

	view.ExecuteTemplate(w, "template", data)
}

func (path Path) Add(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("path/add")...)
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

func (path Path) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	domain := r.FormValue("domain")
	domainID, _ := strconv.Atoi(domain)
	value := r.FormValue("value")
	typeSTR := r.FormValue("type")
	typeID, _ := strconv.Atoi(typeSTR)

	p := &models.Path{
		Name:        name,
		Description: description,
		Domain:      uint(domainID),
		Value:       value,
		Type:        uint(typeID),
	}

	err := dbi.Create(p)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/path", http.StatusSeeOther)
}

func (path Path) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	d, err := dbi.Find([]models.Path{}, map[string]interface{}{"id": ps.ByName("id")})
	if err != nil {
		panic(err)
	}

	dbi.Delete(d)

	http.Redirect(w, r, "/path", http.StatusSeeOther)
}

func (path Path) Update(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	id := Params.ByName("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	domain := r.FormValue("domain")
	domainID, _ := strconv.Atoi(domain)
	value := r.FormValue("value")
	typeSTR := r.FormValue("type")
	typeID, _ := strconv.Atoi(typeSTR)

	p, err := dbi.Find([]models.Path{}, map[string]interface{}{"id": id})

	if err != nil {
		panic(err)
	}

	dbi.Updates(p, models.Path{
		Name:        name,
		Description: description,
		Domain:      uint(domainID),
		Value:       value,
		Type:        uint(typeID),
	})

	http.Redirect(w, r, "/path/update/"+id, http.StatusSeeOther)
}

func (path Path) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("path/get")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Path"], err = dbi.Find([]models.Path{}, map[string]interface{}{"id": ps.ByName("id")})
	if err != nil {
		panic(err)
	}
	data["Domains"], err = dbi.FindAll([]models.Domain{})
	if err != nil {
		panic(err)
	}
	view.ExecuteTemplate(w, "template", data)
}
