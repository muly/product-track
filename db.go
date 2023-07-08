package main

import (
	"context"
	"fmt"
	"net/url"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	tableTrackRequests = "track_requests"
)

type trackInputList []trackInput

func (t *trackInput) id() string {
	return fmt.Sprintf("[%s][%s]", t.TypeOfRequest, url.QueryEscape(t.Url))
}

//get operation using id
func (t *trackInput) getByID(ctx context.Context) error {
	d, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Get(ctx)
	if err != nil {
		return err
	}
	return d.DataTo(t)
}

//delete operation using id 
func (t *trackInput) deleteByID(ctx context.Context) error {
	if _, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Delete(ctx); err != nil {
		return err
	}
	return nil
}

//create operation
func (t *trackInput) create(ctx context.Context) error {
	if _, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Create(ctx, t); err != nil {
		return err
	}
	return nil
}

//update operation
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

//perform set operation
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
		fmt.Println(doc.DataTo(&d))
		*l = append(*l, d)
	}
	return nil
}
