package exporter

import (
	"encoding/csv"
	"fmt"
	customerimporter "importer/importer"
	"io"
)

type CustomerExporter struct {
	output io.Writer
}

// NewCustomerExporter returns a new CustomerExporter that writes customer
// domain data to specified file.
func NewCustomerExporter(output io.Writer) *CustomerExporter {
	return &CustomerExporter{
		output: output,
	}
}

func (ex *CustomerExporter) ExportData(data []customerimporter.DomainData) error {
	if data == nil {
		return fmt.Errorf("data cannot be nil")
	}

	writer := csv.NewWriter(ex.output)
	defer writer.Flush()

	if err := writer.Write([]string{"domain", "count"}); err != nil {
		return fmt.Errorf("failed to write CSV header: %w", err)
	}

	for _, domainData := range data {
		record := []string{
			domainData.Domain,
			fmt.Sprintf("%d", domainData.CustomerQuantity),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write CSV record: %w", err)
		}
	}

	return nil
}
