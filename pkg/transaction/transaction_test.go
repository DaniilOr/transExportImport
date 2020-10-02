package transaction

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestService_MapRowToTransaction(t *testing.T) {
	type fields struct {
		mu           sync.Mutex
		Transactions []*Transaction
	}
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Transaction
		wantErr bool
	}{
		{name: "Empty", fields: fields{mu: sync.Mutex{}, Transactions: []*Transaction{}}, args: args{row: []string{"", "", "", "", "", "",""}}, want: Transaction{ }, wantErr: true},
		{name: "Full", fields: fields{mu: sync.Mutex{}, Transactions: []*Transaction{}}, args: args{row: []string{"1", "123", "321", "552", "ok", "0","10"}}, want: Transaction{1,"123", "321", "552", "ok", time.Unix(0,0), 10}, wantErr: false},
		{name: "No stings", fields: fields{mu: sync.Mutex{}, Transactions: []*Transaction{}}, args: args{row: []string{"1", "", "", "", "", "0","10"}}, want: Transaction{1,"", "", "", "", time.Unix(0,0), 10}, wantErr: false},
		{name: "No ints", fields: fields{mu: sync.Mutex{}, Transactions: []*Transaction{}}, args: args{row: []string{"", "123", "321", "522", "ok", "",""}}, want: Transaction{}, wantErr: true},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				mu:           tt.fields.mu,
				Transactions: tt.fields.Transactions,
			}
			got, err := s.MapRowToTransaction(tt.args.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("MapRowToTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapRowToTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}
