package main

import (
	"fmt"
	"math/rand"
)

var letterBytes string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func main() {
	array := []string{}
	var inputnumber int
	var passwordLength int
	var starElement string
	println("Enter the number of password examples:")
	fmt.Scan(&inputnumber)
	println("Enter the number of characters in the password:")
	fmt.Scan(&passwordLength)
	println("Do you need '* | - | _ | = | ?' sibdols in you password? Y|N")
	fmt.Scan(&starElement)
	if starElement == "N" {
		for i := 0; i < inputnumber; i++ {
			array = append(array, RandStringBytes(passwordLength))
		}
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i], "\n-----------------")
		}
	} else {
		letterBytes += "*-_=?"
		for i := 0; i < inputnumber; i++ {
			array = append(array, RandStringBytes(passwordLength))
		}
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i], "\n-----------------")
		}
	}
	println("✅✅✅✅")

}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
