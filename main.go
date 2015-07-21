package main

import "net/http"

type Scream struct {
	ID       int64
	Username string
	Email    string
	Message  string
	Time     string
}

func init() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/tweets", tweetHandler)
	http.HandleFunc("/profile/", profileHandler)
	http.HandleFunc("/login", loginHandler)
}
