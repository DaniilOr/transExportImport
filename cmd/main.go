package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"github.com/DaniilOr/transExportImport/pkg/transaction"
)

func main() {
	//err := execute("test.csv")
	svc := transaction.NewService()
	err := svc.ImportCSV("test.csv")
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	err = executeExport("test1.csv", svc)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	err = svc.ExportJSON("trans.json")
	if err!= nil{
		os.Exit(1)
	}
	err = svc.ImportJSON("trans.json")
	if err != nil{
		os.Exit(1)
	}
	fmt.Println(svc.Transactions)
}

func executeExport(filename string, svc * transaction.Service) (err error){
	file, err := os.Create(filename)
	if err != nil{
		log.Println(err)
		return err
	}
	defer func(c io.Closer) {
		currentErr := c.Close()
		if currentErr != nil{
			log.Print(currentErr)
			if err == nil{
				err = currentErr
			}
		}
	}(file)
	err = svc.ExportCSV(file)
	if err != nil{
		log.Println(err)
		return err
	}
	return nil
}