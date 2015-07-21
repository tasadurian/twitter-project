package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/login/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(c, "/login/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	getNewTweet(res, req)
	renderHome(res, req)
}

func tweetHandler(res http.ResponseWriter, req *http.Request) {
	displayScreams(res, req)
}

func profileHandler(res http.ResponseWriter, req *http.Request) {
	renderProfile(res, req)
}
