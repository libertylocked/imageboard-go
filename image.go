package vrmp

import (
	"log"
	"time"

	"net/http"
	"net/url"

	"appengine"
	"appengine/blobstore"
	"appengine/datastore"
)

type ImageRecord struct {
	Name         string
	BlobKey      string
	Email        string
	TimeUploaded time.Time
}

func getImageUploadURL(ctx appengine.Context) string {
	uploadURL, err := blobstore.UploadURL(ctx, "/imageupload_complete", nil)
	if err != nil {
		log.Println("Error generating upload URL", err)
		return ""
	}
	return uploadURL.String()
}

func serveImageByKey(w http.ResponseWriter, key string) {
	blobstore.Send(w, appengine.BlobKey(key))
}

func getImageParseUploadKey(ctx appengine.Context, r *http.Request) (string, url.Values) {
	blobs, otherVals, err := blobstore.ParseUpload(r)
	if err != nil {
		log.Println("Error getting blob key!", err)
		return "", otherVals
	}
	file := blobs["file"]
	if len(file) == 0 {
		log.Println("Error! len file is 0", err)
		return "", otherVals
	}
	return string(file[0].BlobKey), otherVals
}

func updateImageRecord(ctx appengine.Context, name, blobkey, email string) {
	img := &ImageRecord{
		Name:         name,
		BlobKey:      blobkey,
		Email:        email,
		TimeUploaded: time.Now(),
	}
	// use blobkey as key
	key := datastore.NewKey(ctx, "ImageRecord", blobkey, 0, nil)
	_, err := datastore.Put(ctx, key, img)
	if err != nil {
		// handle err
		log.Println(err)
		return
	}
}

func getImageRecord(ctx appengine.Context, blobKey string) (ImageRecord, error) {
	key := datastore.NewKey(ctx, "ImageRecord", blobKey, 0, nil)
	var img ImageRecord
	err := datastore.Get(ctx, key, &img)
	if err != nil {
		log.Println(err)
	}
	return img, err
}
