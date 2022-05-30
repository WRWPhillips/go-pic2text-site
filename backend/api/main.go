package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/WRWPhillips/go-pic2text-site/internal"
)

func processFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("./temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)
	fmt.Fprintf(w, "Successfully Uploaded File\n")

	path := fmt.Sprintf("%+v", tempFile.Name())
	width, _ := strconv.Atoi(r.FormValue("width"))
	height, _ := strconv.Atoi(r.FormValue("height"))
	palette := r.FormValue("palette")
	reverse := false

	options := internal.Options{
		Path:    path,
		Width:   width,
		Height:  height,
		Palette: palette,
		Reverse: reverse,
	}

	fmt.Println(&options)
}

func setupRoutes() {
	http.HandleFunc("/upload", processFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Hello World")
	setupRoutes()
}
