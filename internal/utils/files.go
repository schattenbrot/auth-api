package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	fpath "path/filepath"
)

func SaveFiles(filename string, file multipart.File) error {
	workDir, _ := os.Getwd()
	filesDir := fpath.Join(workDir, "data", "files", filename)
	out, err := os.Create(filesDir)
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(out, file)
	return nil
}

func GetFileType(fileHeader *multipart.FileHeader) (string, error) {
	fileType := fpath.Ext(fileHeader.Filename)
	fmt.Println("fileType: ", fileType)
	if fileType == ".jpg" || fileType == ".jpeg" || fileType == ".png" {
		return fileType, nil
	}
	return "", errors.New("file type not supported")
}
