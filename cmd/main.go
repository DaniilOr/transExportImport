package main

import (
	"github.com/DaniilOr/transExportImport/cmd/exporter"
	"github.com/DaniilOr/transExportImport/cmd/importer"
	"github.com/DaniilOr/transExportImport/pkg/transaction"
	"log"
	"os"
)

func main() {

	svc := transaction.NewService()

	err := importer.ExecuteImport("test.csv", svc)

	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	err = exporter.ExecuteExport("test1.csv", svc)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	err = svc.ExportXML("trans.xml")
	if err!= nil{
		os.Exit(1)
	}

	svc = transaction.NewService()

	err = svc.ImportXML("trans.xml")
	if err != nil{
		os.Exit(1)
	}
}

