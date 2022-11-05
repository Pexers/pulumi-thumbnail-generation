package api

import (
	"context"
	"log"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func Data(ctx context.Context, e GCSEvent) error {
	log.Printf("Processing file: %s", e.Name)
	return nil
}
