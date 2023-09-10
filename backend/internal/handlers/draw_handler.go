package handlers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func HandleDrawCommand(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	columnsParam := r.URL.Query().Get("columns")
	columnNames := strings.Split(columnsParam, ",")
	fmt.Println(fileName)
	filePath := "../uploads/" + fileName
	fmt.Println(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to open file: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileReader := csv.NewReader(file)
	fileReader.LazyQuotes = true
	records, err := fileReader.ReadAll()

	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read CSV data: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}
	var extractedColumns []map[string]string
	header := records[0]
	columnIndices := make(map[string]int)
	for i, colName := range header {
		for _, requestedCol := range columnNames {
			if colName == requestedCol {
				columnIndices[colName] = i
			}
		}
	}
	for _, row := range records[1:] {
		extractedRow := make(map[string]string)
		for colName, colIndex := range columnIndices {
			extractedRow[colName] = row[colIndex]
		}
		extractedColumns = append(extractedColumns, extractedRow)
	}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(extractedColumns)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode data as JSON: %s\n", err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
