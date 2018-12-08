package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting extractor")

	fileName := os.Getenv("FILENAME")
	debtType := os.Getenv("DEBT_TYPE")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ex := NewExtractor(debtType)
	if err != nil {
		panic(err.Error)
	}

	for scanner.Scan() {
		line := scanner.Text()
		op, err := ex.Extract(line)
		if err != nil {
			fmt.Println("Error processing line: ", line)
		} else {
			// CALL ELK
			// IF ERROR ENQUEUE
			fmt.Print(op.op)
			fmt.Print(" - ")
			fmt.Println(op.item.data)
		}
	}
}
