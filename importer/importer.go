// Package importer reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type DomainData struct {
	Domain           string
	CustomerQuantity uint64
}

type CustomerImporter struct {
	path string
}

// NewCustomerImporter returns a new CustomerImporter that reads from file at specified path.
func NewCustomerImporter(path string) *CustomerImporter {
	return &CustomerImporter{
		path: path,
	}
}

// ImportDomainData reads and returns sorted customer domain data from CSV file.
func (ci *CustomerImporter) ImportDomainData() ([]DomainData, error) {
	if ci.path == "" {
		return nil, fmt.Errorf("file path cannot be empty")
	}

	file, err := os.Open(ci.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", ci.path, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	csvReader := csv.NewReader(reader)

	header, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	emailColIndex := -1
	for i, col := range header {
		if strings.ToLower(strings.TrimSpace(col)) == "email" {
			emailColIndex = i
			break
		}
	}

	if emailColIndex == -1 {
		return nil, fmt.Errorf("email column not found in CSV header")
	}

	domainCounts := make(map[string]uint64)
	lineNum := 1
	invalidEmailCount := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error reading line %d: %v", lineNum+1, err)
			lineNum++
			continue
		}

		lineNum++

		if len(record) <= emailColIndex {
			log.Printf("warning: line %d has insufficient columns, skipping", lineNum)
			continue
		}

		email := strings.TrimSpace(record[emailColIndex])
		if email == "" {
			continue
		}

		domain, err := extractDomain(email)
		if err != nil {
			log.Printf("warning: invalid email format on line %d: %s, error: %v", lineNum, email, err)
			invalidEmailCount++
			continue
		}

		domainCounts[domain]++
	}

	if invalidEmailCount > 0 {
		log.Printf("warning: skipped %d invalid email addresses", invalidEmailCount)
	}

	result := make([]DomainData, 0, len(domainCounts))
	for domain, count := range domainCounts {
		result = append(result, DomainData{
			Domain:           domain,
			CustomerQuantity: count,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Domain < result[j].Domain
	})

	return result, nil
}

func extractDomain(email string) (string, error) {
	atIndex := strings.LastIndex(email, "@")
	if atIndex == -1 || atIndex == len(email)-1 {
		return "", fmt.Errorf("invalid email format: missing or incomplete domain")
	}

	domain := strings.ToLower(strings.TrimSpace(email[atIndex+1:]))
	if domain == "" {
		return "", fmt.Errorf("invalid email format: empty domain")
	}

	if !strings.Contains(domain, ".") {
		return "", fmt.Errorf("invalid email format: domain must contain a dot")
	}

	return domain, nil
}
