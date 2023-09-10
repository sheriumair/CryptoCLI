package handlers

import (
	"fmt"
	"net/http"
)

/*
*  Handles about endpoint functionality
 */
func HandleAboutCommand(w http.ResponseWriter, r *http.Request) {
	aboutText := "CLI version: 1.0\n"
	aboutText += "This CLI is designed to provide information about cryptocurrency prices.\n"
	aboutText += "Hope you enjoy interacting with it\n"
	aboutText += "Write \" help \" to find more commands\n"

	fmt.Fprint(w, aboutText)
}
