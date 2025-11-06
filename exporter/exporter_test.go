package exporter

import (
	"bytes"
	"importer/importer"
	"os"
	"testing"
)

func TestExportData(t *testing.T) {
	buff := &bytes.Buffer{}
	data := []importer.DomainData{
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

	if len(buff.String()) == 0 {
		t.Fatal("no data written to buffer")
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
	importer := importer.NewCustomerImporter(dataPath)
	data, err := importer.ImportDomainData()
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
