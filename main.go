package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/samuelrms/translate-currency/currency_map"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	inputName := os.Getenv("INPUT_NAME")
	if inputName == "" {
		inputName = "dados.csv"
		fmt.Println("INPUT_NAME not defined; using 'dados.csv'")
	}
	outputName := os.Getenv("OUTPUT_NAME")
	if outputName == "" {
		outputName = "processed.csv"
		fmt.Println("OUTPUT_NAME not defined; using 'processed.csv'")
	}

	// Build paths
	inputPath := filepath.Join("docs", inputName)
	outputDir := "data"
	outputPath := filepath.Join(outputDir, outputName)

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Could not create directory %s: %v\n", outputDir, err)
		os.Exit(1)
	}

	// Open input CSV
	inFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", inputPath, err)
		os.Exit(1)
	}
	defer inFile.Close()

	reader := csv.NewReader(inFile)
	reader.FieldsPerRecord = -1

	// Create output CSV
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", outputPath, err)
		os.Exit(1)
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)

	// Read header
	header, err := reader.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
		os.Exit(1)
	}

	// Locate currency column
	currencyIdx := -1
	for i, col := range header {
		c := strings.ToLower(strings.TrimSpace(col))
		if c == "currency" || c == "moeda" {
			currencyIdx = i
			break
		}
	}

	// If not found, append a new "currency" column
	if currencyIdx < 0 {
		header = append(header, "currency")
		currencyIdx = len(header) - 1
	}

	// Write header
	if err := writer.Write(header); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing header: %v\n", err)
		os.Exit(1)
	}

	// Process rows
	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading record: %v\n", err)
			os.Exit(1)
		}

		// Ensure slice has a spot for currency
		if currencyIdx >= len(rec) {
			rec = append(rec, "")
		}

		// Normalize & map using currency.CurrencyMap
		raw := strings.ToUpper(strings.TrimSpace(rec[currencyIdx]))
		if code, ok := currency_map.CurrencyMap[raw]; ok {
			rec[currencyIdx] = code
		} else {
			rec[currencyIdx] = "BRL"
		}

		// Write out
		if err := writer.Write(rec); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing record: %v\n", err)
			os.Exit(1)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "Error finalizing write: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Currency normalization complete → %s\n", outputPath)
}
