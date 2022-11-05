package api

import (
	"bytes"
	"context"
	"cloud.google.com/go/storage"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"fmt"
)

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func Data(ctx context.Context, e GCSEvent) error {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Printf("Unable to create client")
	}
	// Read image from bucket1
	bucket := client.Bucket(e.Bucket)
	rc, _ := bucket.Object(e.Name).NewReader(ctx)
	defer rc.Close()
	img_bytes, _ := ioutil.ReadAll(rc)

	sub_img_buffer := GenerateThumbnail(img_bytes)

	return nil
}

func GenerateThumbnail(img_bytes []byte) Buffer {
	img, _ := jpeg.Decode(bytes.NewReader(img_bytes))

	sub_img := img.(interface {
        SubImage(r image.Rectangle) image.Image
    }).SubImage(image.Rect(0, 0, 10, 10))

	sub_img_buffer := new(bytes.Buffer)
	err := jpeg.Encode(sub_img_buffer, sub_img, nil)
	if err != nil {
		log.Printf("Error while encoding sub image")
	}

	return sub_img_buffer
}

