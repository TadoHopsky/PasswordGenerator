package main

import (
	"fmt"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_?"

func main() {
	var inputnumber int
	var passwordLength int
	println("Enter --count-- of passwords:")
	fmt.Scan(&inputnumber)
	println("Enter --length-- of password:")
	fmt.Scan(&passwordLength)
	array := []string{}
	for i := 0; i < inputnumber; i++ {
		array = append(array, RandStringBytes(passwordLength))
	}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
