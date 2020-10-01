package transaction

import (
	"encoding/csv"
	"errors"
	"io"
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
	Id int64
	From string
	To string
	MCC string
	Status string
	Date time.Time
	Amount int64
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

func (s * Service) Import(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	s.mu.Lock()
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			if err != nil{
				log.Println(err)
				s.mu.Unlock()
			}
		}
	}(file)
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		s.mu.Unlock()
		log.Println(err)
	}
	s.mu.Unlock()
	for _,row:=range records{
		s.MapRowToTransaction(row)
	}
	if err != nil{
		return err
	}
	return nil
}

func (s * Service) Register(id int64, from string, to string, mcc string, amount int64, status string, date time.Time){
	s.mu.Lock()
	trans := &Transaction{id, from, to, mcc, status, date, amount}
	s.Transactions = append(s.Transactions, trans)
	s.mu.Unlock()
}

func (s * Service) MapRowToTransaction(row[]string){
	if(len(row) != 7){
		log.Println(ErrSpecificationMisspatch)
		os.Exit(1)
	}
	id, err := strconv.ParseInt(row[0], 10, 64)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	date, err := strconv.ParseInt(row[5], 10, 64)
	if err!=nil{
		log.Println(err)
		os.Exit(1)
		}
	amount, err := strconv.ParseInt(row[6], 10, 64)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	s.Register(id, row[1], row[2], row[3], amount, row[4], time.Unix(date, 0))
}

func NewService() *Service{
	return &Service{sync.Mutex{}, []*Transaction{},}
}

