package main

import (
	"context"
	"encoding/json"
	"go-cloudinary/config"
	"go-cloudinary/model"
	"log"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const port = ":8000"

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load file .env : %v", err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Failed to connect database : %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/uploads/image", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		var image model.Image
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		cloudinaryService, _ := cloudinary.NewFromURL(os.Getenv("URL_CLOUDINARY"))
		resp, _ := cloudinaryService.Upload.Upload(ctx, file, uploader.UploadParams{})
		log.Println(resp.SecureURL)

		image.ImageURL = resp.SecureURL
		image.InsertImage(ctx)

		response := struct {
			ImageURL any    `json:"image_url"`
			Message  string `json:"message"`
		}{
			ImageURL: image.ImageURL,
			Message:  "Upload image success",
		}

		json.NewEncoder(w).Encode(response)
	}).Methods(http.MethodPost)

	log.Println("Starting server at localhost" + port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalf("Failed to start server : %v", err)
	}
}
