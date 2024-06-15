package main

import (
	"errors"
	"fmt"
)

func main() {
	//alias INT32 = RUNE | | UINT8 = BYTE

	var number int64 = 1000000
	fmt.Println(number)

	var number2 rune = 1000
	fmt.Println(number2)

	var number3 byte = 100
	fmt.Println(number3)

	var number4 uint32 = 10000
	fmt.Println(number4)

	var number5 float32 = 123.5
	fmt.Println(number5)

	var number6 float64 = 123000.5
	fmt.Println(number6)


	char1 := '$' //36
	char2 := '(' //40
	char3 := 'a' //97
	println(char1, char2, char3)


	var text string = "I eat cupcakes"
	fmt.Println(text)

	text2 := "I can't stop"
	println(text2)



	var boolean bool = false
	println(boolean)

	var boolean2 bool = true
	println(boolean2)


	var erro error = errors.New("Internal error")
	fmt.Println(erro)
}