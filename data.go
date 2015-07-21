package main

import (
	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

func createScream(ctx context.Context, scream *Scream) error {
	u := user.Current(ctx)
	key := datastore.NewKey(ctx, "Scream", u.Email, 0, nil)
	_, err := datastore.Put(ctx, key, scream)
	if err != nil {
		return err
	}
	return nil
}

func getScreams(ctx context.Context) ([]*Scream, error) {
	screams := make([]*Scream, 0)
	q := datastore.NewQuery("Scream")
	iterator := q.Run(ctx)
	for {
		var scream Scream
		key, err := iterator.Next(&scream)
		if err == datastore.Done {
			break
		} else if err != nil {
			return nil, err
		}
		scream.ID = key.IntID()
		screams = append(screams, &scream)
	}
	return screams, nil
}

func createProfile() {

}
