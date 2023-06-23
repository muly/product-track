package main

import (
	"context"
	"fmt"
	"net/url"

	"cloud.google.com/go/firestore"
)

const (
	tableTrackRequests = "track_requests"
)

func (t *trackInput) id() string {
	return fmt.Sprintf("[%s][%s]", url.QueryEscape(t.Url), t.TypeOfRequest)
}

func (t *trackInput) getByID(ctx context.Context) error {
	d, err := client.Collection(tableTrackRequests).Doc(t.id()).Get(ctx)
	if err != nil {
		return err
	}
	if err := d.DataTo(t); err != nil {
		return err
	}
	return nil
}

func (t *trackInput) deleteByID(ctx context.Context) error {
	_, err := client.Collection(tableTrackRequests).Doc(t.id()).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (t *trackInput) create(ctx context.Context) error {
	_, err := client.Collection(tableTrackRequests).Doc(t.id()).Create(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *trackInput) patch(ctx context.Context) error {
	_, err := client.Collection(tableTrackRequests).Doc(t.id()).Update(ctx, []firestore.Update{})
	if err != nil {
		return err
	}
	return nil
}

func (t *trackInput) upsert(ctx context.Context) error {
	_, err := client.Collection(tableTrackRequests).Doc(t.id()).Set(ctx, t)
	if err != nil {
		return err
	}
	return nil
}

// // get using filter:
// // for now the only filter supported is on the active requests field
// func (trackInputRecords) get(ctx context.Context, filters TODO) error {

// 	return errors.New("implementation pending")
// }