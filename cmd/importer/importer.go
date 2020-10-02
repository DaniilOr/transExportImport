package importer

import (
	"github.com/DaniilOr/transExportImport/pkg/transaction"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ExecuteImport(filename string, svc * transaction.Service) (err error){
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
func ExecuteImportJSON(filename string, svc * transaction.Service) (err error){
	file, err := ioutil.ReadFile(filename)
	if err != nil{
		log.Println(err)
		return err
	}
	err = svc.ImportJSON(file)
	if err != nil{
		log.Println(err)
		return err
	}
	return nil
}
