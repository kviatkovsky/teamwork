package exporter

import (
	"bytes"
	customerimporter "importer/importer"
	"os"
	"strings"
	"testing"
)

func TestExportData(t *testing.T) {
	buff := &bytes.Buffer{}
	data := []customerimporter.DomainData{
		{
			Domain:           "livejournal.com",
			CustomerQuantity: 12,
		},
		{
			Domain:           "microsoft.com",
			CustomerQuantity: 22,
		},
		{
			Domain:           "newsvine.com",
			CustomerQuantity: 15,
		},
		{
			Domain:           "pinteres.uk",
			CustomerQuantity: 10,
		},
		{
			Domain:           "yandex.ru",
			CustomerQuantity: 43,
		},
	}
	exporter := NewCustomerExporter(buff)

	err := exporter.ExportData(data)
	if err != nil {
		t.Fatal(err)
	}

	output := buff.String()
	if len(output) == 0 {
		t.Fatal("no data written to buffer")
	}

	if !strings.Contains(output, "domain") || !strings.Contains(output, "count") {
		t.Error("CSV header not found in output")
	}

	if !strings.Contains(output, "livejournal.com") || !strings.Contains(output, "12") {
		t.Error("expected data not found in output")
	}
}

func TestExportEmptyData(t *testing.T) {
	buff := &bytes.Buffer{}
	exporter := NewCustomerExporter(buff)

	err := exporter.ExportData(nil)
	if err == nil {
		t.Fatal(err)
	}
	t.Log(err)
}

func BenchmarkImportDomainData(b *testing.B) {
	b.StopTimer()
	path := "./test_output.csv"

	file, err := os.Open(path)
	if err != nil {
		b.Fatal(err)
	}

	dataPath := "../importer/benchmark10k.csv"
	imp := customerimporter.NewCustomerImporter(dataPath)
	data, err := imp.ImportDomainData()
	if err != nil {
		b.Error(err)
	}
	exporter := NewCustomerExporter(file)

	b.StartTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := exporter.ExportData(data); err != nil {
			b.Fatal(err)
		}
	}
}
