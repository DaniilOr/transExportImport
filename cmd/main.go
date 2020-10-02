package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"github.com/DaniilOr/transExportImport/pkg/transaction"
)

func main() {

	svc := transaction.NewService()

	err := executeImport("test.csv", svc)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(svc.Transactions)
	err = executeExport("test1.csv", svc)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
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
	err = svc.Export(file)
	if err != nil{
		log.Println(err)
		return err
	}
	return nil
}

func executeImport(filename string, svc * transaction.Service) (err error){
	file, err := os.Open(filename)
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

	err = svc.Import(file)
	if err != nil{
		log.Println(err)
		return err
	}
	return nil
}