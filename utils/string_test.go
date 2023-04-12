package utils

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T) {
	var data = "zxcvbnm,"
	bs := Bytes(data)
	fmt.Printf("%s", String(bs))
}
