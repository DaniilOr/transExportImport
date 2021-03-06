package importer

import (
	"github.com/DaniilOr/transExportImport/pkg/transaction"
	"io"
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
