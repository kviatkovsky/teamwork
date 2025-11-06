package customerimporter

import "testing"

func TestImportData(t *testing.T) {
	path := "./test_data.csv"
	importer := NewCustomerImporter(path)

	_, err := importer.ImportDomainData()
	if err != nil {
		t.Error(err)
	}
}

func TestImportDataSort(t *testing.T) {
	sortedDomains := []string{
		"360.cn", "acquirethisname.com", "blogtalkradio.com",
		"chicagotribune.com", "cnet.com", "cyberchimps.com",
		"github.io", "hubpages.com", "rediff.com", "statcounter.com",
	}
	path := "./test_data.csv"
	importer := NewCustomerImporter(path)
	data, err := importer.ImportDomainData()
	if err != nil {
		t.Error(err)
	}
	for i, v := range data {
		if v.Domain != sortedDomains[i] {
			t.Errorf("data not sorted properly. mismatch:\nhave: %v\nwant: %v", v.Domain, sortedDomains[i])
		}
	}
}

func TestImportInvalidPath(t *testing.T) {
	path := ""
	importer := NewCustomerImporter(path)

	_, err := importer.ImportDomainData()
	if err == nil {
		t.Error("invalid path error not caught")
	}
}

func TestImportInvalidData(t *testing.T) {
	path := "./test_invalid_data.csv"
	importer := NewCustomerImporter(path)

	data, err := importer.ImportDomainData()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(data) == 0 {
		t.Error("expected valid emails to be processed, but got no results")
	}

	validDomains := []string{"hubpages.com", "360.cn", "statcounter.com", "prweb.com", "sun.com", "eepurl.com", "domainmarket.com", "unc.edu", "xinhuanet.com"}
	domainMap := make(map[string]bool)
	for _, d := range data {
		domainMap[d.Domain] = true
	}

	for _, domain := range validDomains {
		if !domainMap[domain] {
			t.Errorf("expected domain %s to be present in results", domain)
		}
	}
}

func TestImportReturnsData(t *testing.T) {
	path := "./test_data.csv"
	importer := NewCustomerImporter(path)

	data, err := importer.ImportDomainData()
	if err != nil {
		t.Error(err)
	}
	if len(data) == 0 {
		t.Error("no data returned")
	}
}

func BenchmarkImportDomainData(b *testing.B) {
	b.StopTimer()
	path := "./benchmark10k.csv"
	importer := NewCustomerImporter(path)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := importer.ImportDomainData(); err != nil {
			b.Error(err)
		}
	}
}
