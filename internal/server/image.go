package server 

import (
	"github.com/gin-gonic/gin"
	"github.com/WRWPhillips/go-pic2text-site/internal/utils"
	"github.com/WRWPhillips/go-pic2text-site/internal/store"

	"io/ioutil"
	"fmt"
	"net/http"
	"strconv"
)

func processFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, _, err := ctx.Request.FormFile("myFile")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error(), "location": "formfile"})
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("./internal/store/temp-images", "upload-*.png")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error(), "location": "tempfilestorage"})
		return
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error(), "location": "ioutilreadall"})
		return
	}

	tempFile.Write(fileBytes)

	path := fmt.Sprintf("%+v", tempFile.Name())
	width, _ := strconv.Atoi(ctx.Request.FormValue("width"))
	height, _ := strconv.Atoi(ctx.Request.FormValue("height"))
	palette := ctx.Request.FormValue("palette")
	reverse := false

	options := utils.Options{
		Path:    path,
		Width:   width,
		Height:  height,
		Palette: palette,
		Reverse: reverse,
	}

	
	image := new(store.Image)

	image.Ascii = fmt.Sprintf("%s", &options)

	store.Images = append(store.Images, image)

	ctx.JSON(200, gin.H{
		"msg": "Uploaded successfully",
		"data": image.Ascii,
	})
	return
}
