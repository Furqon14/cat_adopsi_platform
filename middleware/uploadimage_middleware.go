package middleware

import (
	"cat_adoption_platform/model"
	"cat_adoption_platform/service"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type UploadImageMiddleware interface {
	UploadImageMiddleware() gin.HandlerFunc
}

type uploadImageMiddleware struct {
	catService service.CatService
}

func NewUploadImageMiddleware(catService service.CatService) UploadImageMiddleware {
	return &uploadImageMiddleware{catService: catService}
}

func (u *uploadImageMiddleware) UploadImageMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" && strings.Contains(ctx.Request.URL.Path, "/upload-images") {
			// parse multipart fom data
			err := ctx.Request.ParseMultipartForm(10 << 20)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}

			// get all uploaded images
			files := ctx.Request.MultipartForm.File["images"]
			var catImages []model.CatImage
			for _, img := range files {
				img, err := img.Open()
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					ctx.Abort()
					return
				}
				defer img.Close()

				// upload each image to the storage service
				uploadedImage, err := UploadImage(img)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					ctx.Abort()
					return
				}

				// save uploaded image information
				catImage := model.CatImage{
					URL: uploadedImage.URL,
				}
				catImages = append(catImages, catImage)
			}

			// Pass the cat images data to the next handler
			ctx.Set("catImages", catImages)

			ctx.Next()
		} else {
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed"})
			ctx.Abort()
			return
		}
	}
}

func UploadImage(file multipart.File) (model.CatImage, error) {
	// Define the path where the file will be saved
	uploadPath := "./uploads"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.Mkdir(uploadPath, os.ModePerm)
	}
	// Generate a unique filename for the uploaded image
	//  _, filename := filepath.Split(file)
	ext := ".jpg"
	newFilename := fmt.Sprintf("%s%s", "cat-", ext)

	// Save the uploaded file to the specified path
	out, err := os.Create(filepath.Join(uploadPath, newFilename))
	if err != nil {
		return model.CatImage{}, err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return model.CatImage{}, err
	}

	// Return the CatImage struct with the URL
	return model.CatImage{URL: fmt.Sprintf("http://localhost:2000/uploads/%s", newFilename)}, nil
}
