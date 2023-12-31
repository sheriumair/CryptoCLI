package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

/*
*  Handles delete endpoint functionality
 */

func HandleDeleteFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["fileName"]
	filePath := "../uploads/" + fileName

	//Will throw error if file doesnot exists
	err := os.Remove(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to delete file as this file doesnot exist: %s\n", fileName), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("File %s deleted successfully\n", fileName)))
}
