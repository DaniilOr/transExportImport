package exporter

import (
	"encoding/json"
	"github.com/DaniilOr/transExportImport/pkg/transaction"
	"io"
	"io/ioutil"
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
func ExecuteExportJSON(filename string, svc * transaction.Service) (err error){
	encoded, err := json.Marshal(svc.Transactions)
	if err != nil {
		log.Println(err)
		return err
	}
	err = ioutil.WriteFile(filename, encoded, 0644)
	if err != nil{
		log.Println(err)
		return err
	}
	return nil
}