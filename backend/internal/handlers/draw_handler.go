package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
)

func HandleDrawCommand(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	fileName := r.URL.Query().Get("file")
	//columnsParam := r.URL.Query().Get("columns")
	//columnNames := strings.Split(columnsParam, ",")
	fmt.Println(fileName)
	filePath := "../uploads/" + fileName
	fmt.Println(filePath)
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open file: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Parse the CSV data
	fileReader := csv.NewReader(file)
	fileReader.Comma = ',' // Set the comma as the delimiter
	fileReader.LazyQuotes = false
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(records)

	/*// Extract specified columns
	var extractedData []map[string]string
	header := rows[0]
	columnIndices := make(map[string]int)

	for i, colName := range header {
		for _, requestedCol := range columnNames {
			if colName == requestedCol {
				columnIndices[colName] = i
			}
		}
	}

	for _, row := range rows[1:] {
		extractedRow := make(map[string]string)
		for colName, colIndex := range columnIndices {
			extractedRow[colName] = row[colIndex]
		}
		extractedData = append(extractedData, extractedRow)
	}

	// Send extracted data as JSON response
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(extractedData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode data as JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)*/
}
