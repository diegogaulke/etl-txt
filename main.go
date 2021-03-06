package main

import (
	"bufio"
	"etl-txt/extractor"
	"etl-txt/loader"
	"etl-txt/transformer"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting extractor")

	fileName := os.Getenv("FILENAME")
	fileType := os.Getenv("FILE_TYPE")

	e := extractor.NewExtractor(fileType)
	t := transformer.New(fileType)
	p := loader.New(fileType)

	if e == nil || t == nil || p == nil {
		log.Fatal("File type not recognized!")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		op, err := e.Extract(line)
		if err != nil {
			fmt.Println("Error while extracting line: ", line)
		} else {
			err = t.Transform(op)
			if err != nil {
				fmt.Println("Error while transforming line: ", line)
				continue
			}
			err = p.Load(op)
			if err != nil {
				fmt.Println("Error while loading line: ", line)
				continue
			}
		}
	}

	fmt.Println("Done!")
}
