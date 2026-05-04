package xnet

import (
	"net"
	"testing"
)

func TestGetCityByIp(t *testing.T) {
	etIp := "47.96.9.32"
	city := GetCityByIp(etIp)
	t.Log(city)
}

func TestGetCityByIpEmpty(t *testing.T) {
	etIp := ""
	city := GetCityByIp(etIp)
	t.Log(city)
}

func TestGetAreaByIp(t *testing.T) {
	etIp := "47.96.9.32"
	city := GetAreaByIp(etIp)
	t.Log(city.City)
}

func Test_isPrivateSubnet(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test1", args{net.ParseIP("127.0.0.1")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrivateSubnet(tt.args.ip); got != tt.want {
				t.Errorf("isPrivateSubnet() = %v, want %v", got, tt.want)
			}
		})
	}
}
