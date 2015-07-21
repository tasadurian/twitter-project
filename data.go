package main

import (
	"time"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

type Scream struct {
	ID       int64
	Username string
	Email    string
	Message  string
	Time     time.Time
}

func createScream(ctx context.Context, scream *Scream) error {
	//u := user.Current(ctx)
	key := datastore.NewIncompleteKey(ctx, "Scream", nil)
	_, err := datastore.Put(ctx, key, scream)
	if err != nil {
		return err
	}
	return nil
}

func getScreams(ctx context.Context, username string) ([]*Scream, error) {
	screams := make([]*Scream, 0)
	q := datastore.NewQuery("Scream")
	if username != "" {
		q = q.Filter("Username=", username)
	}
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

type Profile struct {
	Username  string
	Email     string
	Following []string
}

func getProfile(ctx context.Context, email string) (*Profile, error) {
	key := datastore.NewKey(ctx, "Profile", email, 0, nil)
	var profile Profile
	return &profile, datastore.Get(ctx, key, &profile)
}

func createProfile(ctx context.Context, profile *Profile) error {
	key := datastore.NewKey(ctx, "Profile", profile.Email, 0, nil)
	_, err := datastore.Put(ctx, key, profile)
	return err
}
func follow(ctx context.Context, email string, followee string) error {
	profile, err := getProfile(ctx, email)
	if err != nil {
		return err
	}
	for _, f := range profile.Following {
		if f == followee {
			return nil
		}
		if followee == "" {
			return nil
		}
	}
	profile.Following = append(profile.Following, followee)
	return createProfile(ctx, profile)
}
