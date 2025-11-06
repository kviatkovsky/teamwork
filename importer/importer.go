// Package importer reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).
package importer

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
func (ci CustomerImporter) ImportDomainData() ([]DomainData, error) {
	return []DomainData{}, nil
}
