package xidcard

import (
	"testing"
)

func TestValidate(t *testing.T) {
	id := "510723198006202551"
	_, err := Validate(id)
	if err != nil {
		t.Error(err)
	}
}

func TestSex(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"t1", args{id: "412924197602102558"}, "男", false},
		{"t2", args{id: "412924197602102558"}, "女", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sex(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sex() = %v, want %v", got, tt.want)
			}
		})
	}
}
