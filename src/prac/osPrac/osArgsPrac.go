package osPrac

import (
	"fmt"
	"os"
)

func OsArgsPrac() {
	//var myString string = "Io Practice"
	var argsWithProg = os.Args
	var argsWithoutProg = os.Args[1:]
	var arg = os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Print(arg)
}