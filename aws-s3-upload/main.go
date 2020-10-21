package main

import (
	"fmt"
	"log"
	"net/http"

	"aws" // AWS SDK
)

func main() {
	AccessKey := ""
	SecretAccessKey := ""
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := awsSession.NewSession(&aws.Config{
		Region: aws.String(""),
		Credentials: credentials.NewStaticCredentials(
			AccessKey,
			SecretAccessKey,
			"",
		),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	fmt.Println("AWS SESSION SUCCESS")
	fmt.Println("CREATING AWS S3 MANAGER UPLOADER")
	// uploader := s3manager.NewUploader(s)

	fmt.Println("READ FORM FILE IMAGE")
	file, handler, err := r.FormFile("image")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//
	defer file.Close()
	fmt.Println("UPLOADING FILE IMAGE")
	fmt.Println(handler.Filename)

	// Upload
	_, errUpload := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(handler.Filename),
		ACL:    aws.String("public-read"),
		Body:   file,
		// ContentDisposition:   aws.String("attachment"),
		// ServerSideEncryption: aws.String("AES256"),
	})

	if errUpload != nil {
		fmt.Println(errUpload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprint(w, "FAILED TO UPLOAD")
		return
	}
	// fmt.Println(result)
	fmt.Println(handler.Filename)

	fmt.Println("UPLOAD PLAYER IMAGE SUCCESS")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprint(w, struct{ Result string }{"https://BUCKET_NAME.s3.REGION.amazonaws.com/" + handler.Filename})
}
