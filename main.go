package main

import (
	"fmt"
	"time"
	"math/rand"
)



func main() {
	//gather characters
	lower := "abcdefghijklmnopqrstuvwxyz"
    upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numbers := "1234567890"
    special := "!@#$%^&*()_+-={}[]|<>/?~`"
    all := lower + upper + numbers + special
	//set length
	length:=6
	//random generator
	source:=rand.NewSource(time.Now().UnixNano())
	random:=rand.New(source)
	password:=""
	for i:=0;i<length;i++{
		password+=string(all[random.Intn(len(all))])
	}
	fmt.Println(password)
}



