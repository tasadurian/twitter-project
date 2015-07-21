package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func getNewTweet(res http.ResponseWriter, req *http.Request) {
	var scream Scream
	json.NewDecoder(req.Body).Decode(&scream)

	ctx := appengine.NewContext(req)
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Scream", u.Email, 0, nil)
	_, err := datastore.Put(ctx, key, &scream)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func displayScreams(res http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	screams := make([]Scream, 0)
	if req.Method == "GET" {
		q := datastore.NewQuery("Scream")
		iterator := q.Run(ctx)
		for {
			var scream Scream
			key, err := iterator.Next(&scream)
			if err == datastore.Done {
				break
			} else if err != nil {
				log.Errorf(ctx, "error retrieving screams: %v", err)
				http.Error(res, err.Error(), 500)
				return
			}
			scream.ID = key.IntID()
			screams = append(screams, scream)
		}
		err := json.NewEncoder(res).Encode(screams)
		if err != nil {
			log.Errorf(ctx, "error marshalling todos: %v", err)
			return
		}
	}
}

func createProfile() {

}
