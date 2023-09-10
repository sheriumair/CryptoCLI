package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
*  Handles upload endpoint functionality
 */
func HandleUploadCommand(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	originalFilename := fileHeader.Filename
	filePath := "../uploads/" + originalFilename
	uploadDir := "../uploads/"

	//Will create a directory if doesnt exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
			return
		}
	}

	destinationFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, file)
	if err != nil {
		http.Error(w, "Failed to write file data", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "File %s uploaded successfully", originalFilename)
}
