package handlerImages

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

	//Get the binary image from  POST request
	picture, _, err := r.FormFile("image")
	//picture, multipartImageHeader, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
	}


	//Create the buffer representing the image header
	imageHeader := make([]byte, 512)
	//Set the buffer content equals to image header
	if _, err := picture.Read(imageHeader); err != nil {
		log.Println(err)
	}
	//Get the content type string
	contentType := http.DetectContentType(imageHeader)
	//Set the buffer back to the start
	if _, err := picture.Seek(0, 0); err != nil {
		log.Println(err)
	}

	//Get the current time
	date := carbon.Now().DateTimeString()
	//Set the folder name equals to the current date (YYYY-MM-DD)
	folder := filepath.Join("static","img", date[:10])
	fmt.Println(folder)
	//Set the file name to the current date (HH:MM:SS)
	fileName := date[11:]

	//If the folder doesn't exist, it creates a new one
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("Creating folder")
		os.MkdirAll(folder, os.ModePerm)
	}

	//Set the full path
	imgPath := filepath.Join(folder, fileName)
	switch contentType {
	case "image/jpeg":
		//Create a file for writing
		file, err := os.Create(imgPath+".jpg")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		//Decode the image
		img, err := jpeg.Decode(picture)
		if err != nil {
			log.Println(err)
		}

		//Save the image to file
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			log.Println(err)
		}
		log.Println("Decoded: %s, MIME Type: %s", imgPath, contentType)

	case "image/png":
		//Create a file for writing
		file, err := os.Create(imgPath+".png")
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		//Decode the image
		img, err := png.Decode(picture)
		if err != nil {
			log.Println(err)
		}

		//Save the image to file
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