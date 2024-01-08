package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	texttemplate "text/template"
	"ulfr/models"

	"github.com/julienschmidt/httprouter"
)

type Fire struct {
	// ...
}

func (fire Fire) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("fire/list")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Fires"], err = dbi.FindAll([]models.Fire{})
	if err != nil {
		panic(err)
	}

	view.ExecuteTemplate(w, "template", data)
}

func (fire Fire) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	f, err := dbi.Find([]models.Fire{}, map[string]interface{}{"random_id": ps.ByName("id")})
	if err != nil {
		panic(err)
	}

	dbi.Delete(f)

	http.Redirect(w, r, "/fire", http.StatusSeeOther)
}

func (fire Fire) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if !SelfControl(r.Host) {
		http.Redirect(w, r, "http://"+r.Host, http.StatusMovedPermanently)
	}

	view, err := template.ParseFiles(Include("fire/get")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Fire"], err = dbi.Find([]models.Fire{}, map[string]interface{}{"random_id": ps.ByName("id")})
	if err != nil {
		panic(err)
	}
	view.ExecuteTemplate(w, "template", data)
}

func (fire Fire) Trigger(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	BrowserInformations := r.FormValue("BrowserInformations")
	UserInformations := r.FormValue("UserInformations")
	Cookies := r.FormValue("Cookies")
	SiteInformations := r.FormValue("SiteInformations")
	Image := r.FormValue("ScreenShot")
	Html := r.FormValue("HTML")
	Host := r.FormValue("host")
	Path := r.FormValue("path")

	randomid := ""
	for {
		randomid, _ = GenerateRandomString(20)
		test, err := dbi.Find([]models.Fire{}, map[string]interface{}{"random_id": randomid})
		if err != nil {
			panic(err)
		}

		if test == nil {
			continue
		} else {
			break
		}
	}

	var browserInformation map[string]interface{}
	err := json.Unmarshal([]byte(BrowserInformations), &browserInformation)
	if err != nil {
		http.Error(w, "JSON çözme hatası: "+err.Error(), http.StatusBadRequest)
		return
	}

	var UserInformation map[string]interface{}
	err = json.Unmarshal([]byte(UserInformations), &UserInformation)
	if err != nil {
		http.Error(w, "JSON çözme hatası: "+err.Error(), http.StatusBadRequest)
		return
	}

	var SiteInformation map[string]interface{}
	err = json.Unmarshal([]byte(SiteInformations), &SiteInformation)
	if err != nil {
		http.Error(w, "JSON çözme hatası: "+err.Error(), http.StatusBadRequest)
		return
	}

	var Cookie map[string]interface{}
	err = json.Unmarshal([]byte(Cookies), &Cookie)
	if err != nil {
		http.Error(w, "JSON çözme hatası: "+err.Error(), http.StatusBadRequest)
		return
	}

	f := &models.Fire{
		URL:                 Host + Path,
		BrowserInformations: browserInformation,
		UserInformations:    UserInformation,
		Cookies:             Cookie,
		SiteInformations:    SiteInformation,
		RandomID:            randomid,
	}

	err = dbi.Create(f)
	if err != nil {
		panic(err)
	}

	WriteToFile("data/"+randomid+"_html.html", Html)
	WriteToFile("data/"+randomid+"_image.txt", Image)
}

func (fire Fire) FireWithPath(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	view, err := texttemplate.ParseFiles(Include("fire/trigger")...)
	if err != nil {
		panic(err)
	}

	data := make(map[string]interface{})
	data["Host"] = r.Host
	data["Path"] = r.RequestURI

	view.ExecuteTemplate(w, "content", data)
}

func (fire Fire) FireWithDomain(w http.ResponseWriter, r *http.Request, Params httprouter.Params) {
	view, err := texttemplate.ParseFiles(Include("fire/trigger")...)
	if err != nil {
		panic(err)
	}
	data := make(map[string]interface{})
	data["Host"] = r.Host
	data["Path"] = r.RequestURI

	view.ExecuteTemplate(w, "content", data)
}

func FireULFR(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	domainURL := strings.Split(host, ":")[0]
	subdomainURL := strings.Split(domainURL, ".")
	obj := new(models.Domain)

	if len(subdomainURL) == 2 || validIP4(strings.Join(subdomainURL, ".")) {
		if err := dbi.Find2(obj, map[string]interface{}{"name": domainURL}); err != nil {
			panic(err)
		}
	} else {
		if err := dbi.Find2(obj, map[string]interface{}{"name": strings.Join(subdomainURL[1:], ".")}); err != nil {
			panic(err)
		}
	}

	PathOBJ := new(models.Path)
	if r.RequestURI == "/" {
		err := dbi.Find2(PathOBJ, map[string]interface{}{
			"domain": obj.ID,
			"value":  subdomainURL[0],
			"type":   "1",
		})
		if err != nil {
			panic(err)
		}
		if PathOBJ.ID != 0 {
			Fire{}.FireWithDomain(w, r, nil)
		}
	} else {
		if obj.Name == domainURL {
			err := dbi.Find2(PathOBJ, map[string]interface{}{
				"domain": obj.ID,
				"value":  strings.Split(r.RequestURI, "/")[1],
				"type":   "2",
			})
			if err != nil {
				panic(err)
			}
			if PathOBJ.ID != 0 {
				Fire{}.FireWithPath(w, r, nil)
			}
		}
	}
}
