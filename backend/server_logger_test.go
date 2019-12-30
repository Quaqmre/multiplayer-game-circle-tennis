package backend

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestLogOutput(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := NewServerLogger(buffer)
	logger.LogStatus("deneme")
	index, count := 0, 0
	array := buffer.Bytes()
	for c, i := range array {
		if i == byte('[') {
			count++
			if count == 2 {
				index = c
				break
			}
		}
	}
	str := array[index:]
	fmt.Println(strings.Trim(string(str), " "))
	if string(str) != "[  deneme  ]\n" {
		t.Error("LogStatus output syntax it shoul changed")
	}
}
