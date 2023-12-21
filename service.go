package main

import (
	"context"
	"log"
	"net/url"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recordUnsupportedWebsiteVisits(ctx context.Context, urlString string) error {
	u, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	t := &unsupportedWebsiteVisits{host: u.Hostname()}

	if err := t.getByID(ctx); err != nil {
		if status.Code(err) != codes.NotFound {
			return err
		}
	}
	t.visitCount = t.visitCount + 1

	if err := t.upsert(ctx); err != nil {
		return err
	}

	log.Printf("successfully recorded the unsupported website: %s", u.Hostname())
	return nil
}
