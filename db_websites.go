package main

import (
	"context"
	"log"
)

// visited unsupported websites

const tableUnsupportedWebsiteVisits = "unsupported_website"

func (t *unsupportedWebsiteVisits) id() string {
	return t.host
}

func (t *unsupportedWebsiteVisits) getByID(ctx context.Context) error {
	d, err := firestoreClient.Collection(tableUnsupportedWebsiteVisits).Doc(t.id()).Get(ctx)
	if err != nil {
		log.Println("(t *unsupportedWebsiteVisits) getByID() error", err)
		return err
	}
	return d.DataTo(t)
}

func (t *unsupportedWebsiteVisits) upsert(ctx context.Context) error {
	if _, err := firestoreClient.Collection(tableUnsupportedWebsiteVisits).Doc(t.id()).Set(ctx, t); err != nil {
		return err
	}
	return nil
}
