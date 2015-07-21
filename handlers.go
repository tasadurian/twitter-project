package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
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

func profileHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "templates/profile.html")
}
func homeHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "templates/home.html")
}

func tweetHandler(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	switch req.Method {
	case "GET":
		// get a list of tweetss
		screams, err := getScreams(ctx)
		if err != nil {
			log.Errorf(ctx, "error getting screams: %v", err)
			return
		}
		err = json.NewEncoder(res).Encode(screams)
		if err != nil {
			log.Errorf(ctx, "error marshalling todos: %v", err)
			return
		}

	case "POST":
		// create a tweet
		var scream Scream
		err := json.NewDecoder(req.Body).Decode(&scream)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		err = createScream(ctx, &scream)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		json.NewEncoder(res).Encode(&scream)

	case "DELETE":
		// delete a tweet
	default:
		http.Error(res, "method not allowed", 405)
	}
}
