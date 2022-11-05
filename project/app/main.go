package api

import (
	"bytes"
	"context"
	"cloud.google.com/go/storage"
	"image"
	"image/jpeg"
	"io"
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
	bucket1 := client.Bucket(e.Bucket)
	rc, _ := bucket1.Object(e.Name).NewReader(ctx)
	defer rc.Close()
	img_bytes, _ := ioutil.ReadAll(rc)

	sub_img_buffer := GenerateThumbnail(img_bytes)

	// Upload thumbnail to bucket2
	object := "thumbnail-" + e.Name
	bucket2 := client.Bucket("bucket2thumbnails")
	wc := bucket2.Object(object).NewWriter(ctx)
	wc.ChunkSize = 0 // note retries are not supported for chunk size 0.
	if _, err = io.Copy(wc, sub_img_buffer); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	log.Printf(object + " uploaded to bucket2.")

	return nil
}

func GenerateThumbnail(img_bytes []byte) *bytes.Buffer {
	img, _ := jpeg.Decode(bytes.NewReader(img_bytes))
	img_bounds := img.Bounds()

	sub_img := img.(interface {
        SubImage(r image.Rectangle) image.Image
    }).SubImage(image.Rect(0, 0, img_bounds.Dx()/2, img_bounds.Dy()))

	sub_img_buffer := new(bytes.Buffer)
	err := jpeg.Encode(sub_img_buffer, sub_img, nil)
	if err != nil {
		log.Printf("Error while encoding sub image")
	}

	return sub_img_buffer
}
