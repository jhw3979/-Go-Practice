package ioPrac

import (
	"io"
	"os"
)

func IoPrac() {
	var myString string = ""
	var args = os.Args

	if len(args) == 1 {
		myString = "인자를 입력하세요"
	} else {
		myString = args[1]
	}

	io.WriteString(os.Stdout, myString)
}