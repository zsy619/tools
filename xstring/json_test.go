package xstring

import "testing"

func Test_JsonToObject(t *testing.T) {
	type args struct {
		meta   string
		result any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := JsonToObject(tt.args.meta, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("JsonToObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
