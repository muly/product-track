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
)

type trackInputList []trackInput

// patchList struct defines the patch payload for multiple patches together
type patch struct {
	typeOfRequest string
	url           string
	patchData     map[string]interface{}
}

type patchList []patch

func (t *trackInput) id() string {
	return fmt.Sprintf("[%s][%s]", t.TypeOfRequest, url.QueryEscape(t.Url))
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
			Url:           p.url,
			TypeOfRequest: p.typeOfRequest,
		}
		if err := t.patch(ctx, p.patchData); err != nil {
			log.Printf("Failed to update process fields for id %s: %v", t.id(), err)
			// Note: no need to return error, just continue processing next record
			continue
		}
	}
}
