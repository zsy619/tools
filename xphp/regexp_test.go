package xphp

import (
	"fmt"
	"testing"
)

func TestPregMatch(t *testing.T) {
	str := `100.100.100.100 - - [23/Feb/2015:03:03:56 +0100] "GET /folder/file.mp3 HTTP/1.1" 206 5637064 "-" "AppleCoreMedia/1.0.0.12B466 (iPhone; U; CPU OS 8_1_3 like Mac OS X; da_dk)"`
	math := `([.0-9]+) .*?\[([0-9a-zA-Z:\/+ ]+)\].*?"[A-Z]+ \/([^\/ ]+)\/([a-zA-Z0-9\-.]+).*" ([0-9]{3}) .*"(.*?)"$`
	strs, err := PregMatch(math, str)
	if err != nil {
	} else {
		for i, s := range strs {
			fmt.Println("", i, s)
		}
	}
}
