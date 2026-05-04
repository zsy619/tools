package xcache

import (
	"fmt"
	"testing"
)

func Test_Diskv(t *testing.T) {
	key := "alpha"
	diskKeyValue.Write(key, []byte{'1', '2', '3'})

	key1 := "alphax"
	diskKeyValue.Write(key1, []byte{'1', '2', '3', '4'})

	// Read the value back out of the store.
	value, _ := diskKeyValue.Read(key)
	fmt.Printf("%v\n", value)

	// Erase the key+value from the store (and the disk).
	// diskKeyValue.Erase(key)
}

func TestSetDisCahce(t *testing.T) {
	SetDisCahce("abc", []byte("sss"))
}

func TestGetDisCahce(t *testing.T) {
	fmt.Println(DisCahceExists("abcx"))
	find, err := GetDisCahce("abc")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(find)
}
