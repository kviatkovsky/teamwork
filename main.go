package main

import (
	"flag"
	"fmt"
	"importer/exporter"
	customerimporter "importer/importer"
	"log"
	"os"
)

func main() {
	var (
		inputFile  = flag.String("input", "customers.csv", "Path to input CSV file")
		outputFile = flag.String("output", "", "Path to output CSV file (if empty, outputs to terminal)")
	)
	flag.Parse()

	if *inputFile == "" {
		log.Fatal("error: input file path cannot be empty")
	}

	log.Printf("Reading customer data from %s...", *inputFile)
	customerImporter := customerimporter.NewCustomerImporter(*inputFile)
	domainData, err := customerImporter.ImportDomainData()
	if err != nil {
		log.Fatalf("error importing domain data: %v", err)
	}

	log.Printf("Found %d unique email domains", len(domainData))

	var outputWriter *os.File
	if *outputFile != "" {
		outputWriter, err = os.Create(*outputFile)
		if err != nil {
			log.Fatalf("error creating output file %s: %v", *outputFile, err)
		}
		defer outputWriter.Close()
		log.Printf("Writing results to %s...", *outputFile)
	} else {
		outputWriter = os.Stdout
	}

	exporter := exporter.NewCustomerExporter(outputWriter)
	if err := exporter.ExportData(domainData); err != nil {
		log.Fatalf("error exporting data: %v", err)
	}

	if *outputFile != "" {
		log.Printf("Successfully wrote results to %s", *outputFile)
	} else {
		fmt.Println()
	}
}
