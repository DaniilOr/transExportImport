package exporter

import (
	"github.com/DaniilOr/transExportImport/pkg/transaction"
	"io"
	"log"
	"os"
)

func ExecuteExport(filename string, svc * transaction.Service) (err error){
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

