package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func Home(w http.ResponseWriter, r *http.Request) {

	renderPages(w, "home.jet", nil)

}

func renderPages(w http.ResponseWriter, tmpl string, data jet.VarMap) {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
