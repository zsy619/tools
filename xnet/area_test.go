package xnet

import (
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
