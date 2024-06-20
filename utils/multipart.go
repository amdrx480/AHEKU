package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// SaveUploadedFile saves a multipart file to the specified directory and returns the file path
func SaveUploadedFile(file *multipart.FileHeader, uploadDir string) (string, error) {
	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create the directory if it doesn't exist
	// err = os.MkdirAll(uploadDir, os.ModePerm)
	// if err != nil {
	// 	return "", err
	// }

	// Buat nama file unik
	fileName := fmt.Sprintf("%s_%d%s", strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename)), time.Now().UnixNano(), filepath.Ext(file.Filename))
	filePath := filepath.Join(uploadDir, fileName)

	// Create the file path
	// filePath := filepath.Join(uploadDir, file.Filename)

	// Create the destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file contents to the destination file
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	// Return the file path
	return filePath, nil
}

func GetFileName(filePath string) string {
	return filepath.Base(filePath)
}
