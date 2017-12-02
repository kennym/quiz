package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open("problems.csv")
	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	// Automatically close file after running program
	defer file.Close()

	reader := csv.NewReader(file)
	totalCount := 0
	correctCount := 0

	for {
		// Read just one record, but we could ReadAll() as well
		record, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("Error:", error)
			return
		}

		fmt.Print(" ", record[0], " = ")

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == record[1] {
				correctCount++
			}
			break
		}

		totalCount += 1
	}

	fmt.Println("Correct results: ", correctCount)
	fmt.Println("Total results: ", totalCount)
}
