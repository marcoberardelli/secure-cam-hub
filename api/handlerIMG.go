package main

import (
	"image/jpeg"
	"log"
	"os"
	"image/png"
	"net/http"
	"github.com/uniplaces/carbon"
	"path/filepath"
	"fmt"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func PostImage(w http.ResponseWriter, r *http.Request) {



	r.ParseForm()
	//key:= r.FormValue("Key")
	//key := r.Form.Get("Key")
	key := r.PostForm.Get("Key")
	fmt.Println(key)


	// Gets the binary image from  POST request
	picture, _, err := r.FormFile("image")
	//picture, multipartImageHeader, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
	}


	// Creates the buffer representing the image header
	imageHeader := make([]byte, 512)
	//Set the buffer content equals to image header
	if _, err := picture.Read(imageHeader); err != nil {
		log.Println(err)
	}
	// Gets the content type string
	contentType := http.DetectContentType(imageHeader)
	// Sets the buffer back to the start
	if _, err := picture.Seek(0, 0); err != nil {
		log.Println(err)
	}

	// Gets the current time
	date := carbon.Now().DateTimeString()
	// Sets the folder name equals to the current date (YYYY-MM-DD)
	folder := filepath.Join("static","img", date[:10])
	fmt.Println(folder)
	// Sets the file name to the current date (HH:MM:SS)
	fileName := date[11:]

	// If the folder doesn't exist, it creates a new one
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("Creating folder")
		os.MkdirAll(folder, os.ModePerm)
	}

	// Sets the full path
	imgPath := filepath.Join(folder, fileName)
	switch contentType {
	case "image/jpeg":
		// Creates a file for writing
		file, err := os.Create(imgPath+".jpg")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		// Decodes the image
		img, err := jpeg.Decode(picture)
		if err != nil {
			log.Println(err)
		}

		// Saves the image to file
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("Decoded: %s, MIME Type: %s", imgPath, contentType)

	case "image/png":
		// Creates a file for writing
		file, err := os.Create(imgPath+".png")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		// Decodes the image
		img, err := png.Decode(picture)
		if err != nil {
			log.Println(err)
		}

		// Save the images to file
		err = png.Encode(file, img)
		if err != nil {
			log.Println(err)
		}
		log.Println("Decoded: %s, MIME Type: %s", imgPath, contentType)

	default:
		log.Println("Unknown MIME Type")
	}
	//fmt.Println("Success!")
}