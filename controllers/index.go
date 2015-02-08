package controllers

import (
	"../utilities"
	"encoding/json"
	"net/http"
)

type Index struct{}
type SiteMap []map[string]string

func (i *Index) Welcome(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "static/index.html")
}

func (i *Index) Sitemap(w http.ResponseWriter, req *http.Request) {
	sitemap := SiteMap{
		{"url": "#/login",
			"title": "Login"},
	}
	sitemap_for_user := SiteMap{
		{"url": "/auth/logout",
			"title": "Logout"},
		{"url": "#/users/profile",
			"title": "Profile"},
	}
	_, err := utilities.GetUserId(req)
	if err == nil {
		js, _ := json.Marshal(sitemap_for_user)
		w.Write(js)
	} else {
		js, _ := json.Marshal(sitemap)
		w.Write(js)
	}
}
