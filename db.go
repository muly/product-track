package main

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

const (
	tableTrackRequests = "track_requests"
)

type trackInputList []trackInput

func (t *trackInput) id() string {
	return fmt.Sprintf("[%s][%s]", url.QueryEscape(t.Url), t.TypeOfRequest)
}

func (t *trackInput) getByID(ctx context.Context) error {
	d, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Get(ctx)
	if err != nil {
		return err
	}
	if err := d.DataTo(t); err != nil {
		return err
	}
	return nil
}

func (t *trackInput) deleteByID(ctx context.Context) error {
	_, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (t *trackInput) create(ctx context.Context) error {
	_, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Create(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *trackInput) patch(ctx context.Context) error {
	currentTime := time.Now()

	_, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Update(ctx, []firestore.Update{
		{Path: "/execute-request", Value: currentTime},
		{Path: "/excute-request", Value: "SUCCESS"},
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *trackInput) upsert(ctx context.Context) error {
	_, err := firestoreClient.Collection(tableTrackRequests).Doc(t.id()).Set(ctx, t)
	if err != nil {
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
