package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)


func getRandomPoem() string {
	// Load the poems csv
	file, err := os.Open("poems.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	total_poems := len(records) - 1           // Subtracting the header
	random_poem := rand.Intn(total_poems) + 1 // Add 1 to avoid header
	return records[random_poem][1]
}
