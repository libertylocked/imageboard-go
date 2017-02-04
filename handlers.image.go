package vrmp

import (
	"io"
	"log"
	"net/http"
	"time"
)

func handleImageUpload(w http.ResponseWriter, r *http.Request) {
	postURL := getImageUploadURL(getContext(r))
	tmplData := map[string]string{
		"uploadURL": postURL,
	}
	renderTemplate(w, "image_upload.html", tmplData)
}

func handleImageUploadComplete(w http.ResponseWriter, r *http.Request) {
	ctx := getContext(r)
	// XXX: blobstore parse must be called before r.FormValue
	blobKey, otherValues := getImageParseUploadKey(ctx, r)
	name := otherValues.Get("name")
	log.Println("Storing", name, " blob key:", blobKey)
	if blobKey != "" {
		updateImageRecord(ctx, name, blobKey, getUserEmail(r))
		http.Redirect(w, r, "/imageview?blobkey="+blobKey, http.StatusFound)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func handleImageView(w http.ResponseWriter, r *http.Request) {
	blobKey := r.FormValue("blobkey")
	if blobKey == "" {
		return
	}
	img, err := getImageRecord(getContext(r), blobKey)
	if err != nil {
		io.WriteString(w, "image record not found")
		return
	}
	tmplData := map[string]string{
		"name":         img.Name,
		"blobKey":      img.BlobKey,
		"email":        img.Email,
		"timeUploaded": img.TimeUploaded.Format(time.UnixDate),
	}
	renderTemplate(w, "image_view.html", tmplData)
}

func handleImageServe(w http.ResponseWriter, r *http.Request) {
	blobKey := r.FormValue("blobkey")
	if blobKey != "" {
		serveImageByKey(w, blobKey)
	}
}
