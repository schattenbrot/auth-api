package utils

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

func GetFile(filename string) ([]byte, error) {
	workDir, _ := os.Getwd()
	filesDir := fpath.Join(workDir, "data", "files", filename)
	file, err := os.Open(filesDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}
	return bytes, nil
}

func GetFileType(fileHeader *multipart.FileHeader) (string, error) {
	fileType := fpath.Ext(fileHeader.Filename)
	fmt.Println("fileType: ", fileType)
	if fileType == ".jpg" || fileType == ".jpeg" || fileType == ".png" {
		return fileType, nil
	}
	return "", errors.New("file type not supported")
}
