package storage

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func UploadImage(image *multipart.FileHeader) (filename string, err error) {
	config := &firebase.Config{
		StorageBucket: os.Getenv("FIREBASE_STORAGE_BUCKET_URL"),
	}

	resp, err := http.Get(os.Getenv("FIREBASE_SECRETS_FILE_URL"))
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}

	opt := option.WithCredentialsJSON(data)
	app, err := firebase.NewApp(context.Background(), config, opt)
	ctx := context.Background()

	if err != nil {
		log.Fatalln(err)
		return
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
		return
	}

	src, err := image.Open()

	if err != nil {
		return
	}
	defer src.Close()

	filename = strings.Trim(strings.ToLower(image.Filename), " ")

	contentType := image.Header.Get("Content-Type")

	writer := bucket.Object("images/" + filename).NewWriter(ctx)
	writer.ObjectAttrs.ContentType = contentType
	writer.ObjectAttrs.CacheControl = "no-cache"

	if _, err = io.Copy(writer, src); err != nil {
		log.Fatalln(err)
		return
	}

	if err = writer.Close(); err != nil {
		log.Fatalln(err)
		return
	}

	return
}
