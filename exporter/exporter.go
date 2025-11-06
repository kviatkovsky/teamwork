package exporter

import (
	"importer/importer"
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

// ExportData writes sorted customer domain data to a CSV file. If file already
// exists, it should be truncated.
func (ex CustomerExporter) ExportData(data []importer.DomainData) error {
	return nil
}
