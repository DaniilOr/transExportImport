package transaction

import (
	"encoding/csv"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)
var (
	ErrSpecificationMisspatch = errors.New("Number of attributes mismatch")
)

type Transaction struct{
	Id int64 `xml:"id"`
	From string `xml:"from"`
	To string `xml:"to"`
	MCC string `xml:"mcc"`
	Status string `xml:"status"`
	Date time.Time `xml:"date"`
	Amount int64 `xml:"amount"`
}
type Service struct{
	mu sync.Mutex
	Transactions []*Transaction
}

func (s * Service) Export(writer io.Writer) error{
	s.mu.Lock()
	if len(s.Transactions) == 0{
		s.mu.Unlock()
		return nil
	}
	records := make([][]string, 0)
	for _, t := range s.Transactions{
		record := []string{
			strconv.FormatInt(t.Id, 10),
			t.From,
			t.To,
			t.MCC,
			t.Status,
			strconv.FormatInt(t.Date.Unix(), 10),
			strconv.FormatInt(t.Amount, 10),
		}
		records = append(records, record)
	}
	s.mu.Unlock()
	w := csv.NewWriter(writer)
	return w.WriteAll(records)
}

func (s * Service) Import(file io.Reader) (err error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		return err
	}
	for _,row:=range records{
		transaction, err := s.MapRowToTransaction(row)
		if err != nil{
			return  err
		}
		s.Register(transaction)
	}
	if err != nil{
		return err
	}
	return nil
}

func (s * Service) Register(transaction Transaction){
	s.mu.Lock()
	s.Transactions = append(s.Transactions, &transaction)
	s.mu.Unlock()
}

func (s * Service) MapRowToTransaction(row[]string) (Transaction, error){
	if(len(row) != 7){
		log.Println(ErrSpecificationMisspatch)
		return Transaction{}, ErrSpecificationMisspatch
	}
	id, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		log.Println(err)
		return Transaction{}, err
	}
	date, err := strconv.ParseInt(row[5], 10, 64)
	if err!=nil{
		log.Println(err)
		return Transaction{}, err
		}
	amount, err := strconv.ParseInt(row[6], 10, 64)
	if err != nil{
		log.Println(err)
		return Transaction{}, err
	}
	return Transaction{id, row[1], row[2], row[3], row[4],time.Unix(date, 0),  amount,}, nil

}

func (s*Service) ImportXML(file string) error{
	data, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return err
	}
	defer data.Close()
	decoder := xml.NewDecoder(data)
	for {
		tok, err := decoder.Token()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "Transaction" {
				var transaction Transaction
				decoder.DecodeElement(&transaction, &tp)
				s.Transactions = append(s.Transactions, &transaction)
			}
		}
	}
	return nil
}
func (s*Service) ExportXML(file string) error{

	encoded, err := xml.Marshal(s.Transactions)
	if err != nil {
		log.Println(err)
		return err
	}
	encoded = append([]byte(xml.Header), encoded...)
	err = ioutil.WriteFile(file, encoded, 0644)
	if err != nil{
		log.Println(err)
		return err
	}
	return nil
}
func NewService() *Service{
	return &Service{sync.Mutex{}, []*Transaction{},}
}