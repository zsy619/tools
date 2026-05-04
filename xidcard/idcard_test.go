package xidcard

import (
	"testing"
)

func TestIdCard_Birth(t *testing.T) {
	tests := []struct {
		name    string
		id      IdCard
		want    string
		wantErr bool
	}{
		{"t1", NewIdCard("412924197602102558"), "19760210", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.id.Birth()
			if (err != nil) != tt.wantErr {
				t.Errorf("IdCard.Birth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IdCard.Birth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdCard_BirthYm(t *testing.T) {
	tests := []struct {
		name    string
		id      IdCard
		want    string
		wantErr bool
	}{
		{"t1", NewIdCard("412924197602102558"), "197602", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.id.BirthYm()
			if (err != nil) != tt.wantErr {
				t.Errorf("IdCard.BirthYm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IdCard.BirthYm() = %v, want %v", got, tt.want)
			}
		})
	}
}
