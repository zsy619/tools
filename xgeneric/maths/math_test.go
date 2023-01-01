package maths

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		n uint
		m uint
	}
	tests := []struct {
		name  string
		args  args
		want  uint
		want1 uint
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{n: 10, m: 3}, want: 4, want1: 2},
		{name: "test2", args: args{n: 10, m: 10}, want: 10, want1: 10},
		{name: "test3", args: args{n: 10, m: 9}, want: 9, want1: 9},
		// {name: "test4", args: args{n: 10, m: 0}, want: 0, want1: 0},
		// {name: "test5", args: args{n: 0, m: 0}, want: 0, want1: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Split(tt.args.n, tt.args.m)
			if got != tt.want {
				t.Errorf("Split() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Split() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{a: 1, b: 2}, want: 2},
		{name: "test2", args: args{a: 2, b: 1}, want: 2},
		{name: "test3", args: args{a: 2, b: 2}, want: 2},
		{name: "test4", args: args{a: 3, b: 4}, want: 4},
		{name: "test5", args: args{a: 0, b: 4}, want: 4},
		{name: "test6", args: args{a: 0, b: 0}, want: 0},
		{name: "test7", args: args{a: 1, b: -1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
