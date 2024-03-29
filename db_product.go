package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	tableTrackRequests = "track_requests"
	tableUsers         = "user"
)

type patch struct {
	// key fields
	typeOfRequest string
	emailID       string
	url           string
	zipCode       int
	// data fields in map
	patchData map[string]interface{}
}

// patchList struct defines the patch payload for multiple patches together
type patchList []patch

type trackInputList []trackInput

func (t *trackInput) id() string {
	return fmt.Sprintf("[%s][%s][%s]", t.TypeOfRequest, t.EmailID, url.QueryEscape(t.URL))
}

// get operation using id
func (t *trackInput) getByID(ctx context.Context) error {
	d, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Get(ctx)
	if err != nil {
		return err
	}
	return d.DataTo(t)
}

// delete operation using id
func (t *trackInput) deleteByID(ctx context.Context) error {
	if _, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Delete(ctx); err != nil {
		return err
	}
	return nil
}

// create operation
func (t *trackInput) create(ctx context.Context) error {
	if _, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Create(ctx, t); err != nil {
		return err
	}
	return nil
}

// update operation
func (t *trackInput) patch(ctx context.Context, updates map[string]interface{}) error {
	// convert the map to slice
	firestoreUpdate := make([]firestore.Update, 0, len(updates))
	for key, value := range updates {
		firestoreUpdate = append(firestoreUpdate, firestore.Update{
			Path:  key,
			Value: value,
		})
	}
	if _, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Update(ctx, firestoreUpdate); err != nil {
		return err
	}
	return nil
}

// perform set operation
func (t *trackInput) upsert(ctx context.Context) error {
	if _, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Set(ctx, t); err != nil {
		return err
	}
	return nil
}

type filter struct {
	path  string
	op    string
	value interface{}
}

// get using filter:
func (l *trackInputList) get(ctx context.Context, filters []filter) error {
	q := firestoreClient.Collection(tableTrackRequests).Query
	for _, f := range filters {
		q = q.Where(f.path, f.op, f.value)
	}
	iter := q.Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		d := trackInput{}
		if err := doc.DataTo(&d); err != nil {
			return err
		}
		*l = append(*l, d)
	}
	return nil
}

// patch runs the patch updates for the given multiple patch records ignoring the db errors
func (pl patchList) patch(ctx context.Context) {
	for _, p := range pl {
		t := trackInput{
			URL:           p.url,
			EmailID:       p.emailID,
			TypeOfRequest: p.typeOfRequest,
			ZipCode:       p.zipCode,
		}
		if err := t.patch(ctx, p.patchData); err != nil {
			log.Printf("Failed to update process fields for id %s: %v", t.id(), err)
			// Note: no need to return error, just continue processing next record
			continue
		}
	}
}

func (u *User) upsert(ctx context.Context) error {
	log.Println("storing email is started")
	_, err := firestoreClient.Collection(tableUsers).Doc(u.Email).Set(ctx, map[string]interface{}{})
	if err != nil {
		log.Printf("Failed to store email in Firestore: %v", err)
		return err
	}
	return nil
}
