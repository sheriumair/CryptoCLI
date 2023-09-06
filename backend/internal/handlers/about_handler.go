package handlers

import (
	"fmt"
	"net/http"
)

func HandleAboutCommand(w http.ResponseWriter, r *http.Request) {
	aboutText := "CLI version: 1.0\n"
	aboutText += "This CLI is designed to provide information about cryptocurrency prices.\n"

	fmt.Fprint(w, aboutText)
}
