package main

import (
	"context"
	"errors"
	"fmt"
	"net/url"
)

const (
	tableTrackRequests = "track_requests"
)

// - get using filter
// - get by id
// - delete by id
// - create/insert
// - update by id

// type trackRequest struct{}

type trackInputRecords []trackInput

func (t *trackInput) id() string {
	return fmt.Sprintf("%s|%s", url.QueryEscape(t.Url), t.TypeOfRequest)
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

	return errors.New("implementation pending")
}

func (t *trackInput) create(ctx context.Context) error {

	return errors.New("implementation pending")
}

func (t *trackInput) update(ctx context.Context) error {

	return errors.New("implementation pending")
}

// // get using filter:
// // for now the only filter supported is on the active requests field
// func (trackInputRecords) get(ctx context.Context, filters TODO) error {

// 	return errors.New("implementation pending")
// }
