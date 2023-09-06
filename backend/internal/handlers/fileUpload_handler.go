package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleUploadCommand(w http.ResponseWriter, r *http.Request) {
	// Get the file from the request
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	originalFilename := fileHeader.Filename
	// Define the destination file path
	filePath := "../uploads/" + originalFilename

	// Create the destination file
	destinationFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer destinationFile.Close()

	// Write the binary data to the destination file
	_, err = io.Copy(destinationFile, file)
	if err != nil {
		http.Error(w, "Failed to write file data", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	fmt.Fprintf(w, "File %s uploaded successfully", originalFilename)
}
