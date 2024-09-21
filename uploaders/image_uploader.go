// uploaders/image_uploader.go
package uploaders

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type LocalImageUploader struct {
	BasePath string
}

func (uploader *LocalImageUploader) Upload(file multipart.File, filename string) (string, error) {
	// Implement the logic to save the file locally
	// For simplicity, let's assume we save it to the BasePath
	path := fmt.Sprintf("%s/%s", uploader.BasePath, filename)
	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return path, nil
}

func NewImageUploader(BasePath string) LocalImageUploader {
	return LocalImageUploader{BasePath: BasePath}
}
