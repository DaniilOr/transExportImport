package main

import (
	"io"
	"log"
	"os"
	"transExportImport/pkg/transaction"
)

func main() {
	//err := execute("test.csv")
	svc := transaction.NewService()
	svc.Import("test.csv")
	err := execute("test1.csv", svc)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
}

func execute(filename string, svc * transaction.Service) (err error){
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