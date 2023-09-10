package handlers

import (
	"fmt"
	"net/http"
)

func HandleHelpCommand(w http.ResponseWriter, r *http.Request) {
	helpText := "Available commands:\n"
	helpText += "help - Displays a list of all available commands and briefly describes what they do.\n"
	helpText += "about - Shows information about the CLI, including its version and the purpose of the project.\n"
	helpText += "fetch-price/{pair} - Fetches the current average price of a specified cryptocurrency pair e.g fetch-price BTCUSDT.\n"
	helpText += "upload - Uploads the file in our backend.\n"
	helpText += "draw [Filename.ext] [columns] - Draw the specified columns of the given file e.g draw SolanaHistoricalData.csv Price,Open .\n"
	helpText += "delete [Filename.ext] - Delete the file from our backend e.g delete SolanaHistoricalData.csv .\n"

	fmt.Fprint(w, helpText)
}
