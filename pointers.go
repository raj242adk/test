package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("my stirng is set to",myString)
	changeUsingPointer(&myString)
	log.Println("my stirng is set to",myString)

}

func changeUsingPointer(s *string){
	newvalue :="Red"
	*s =newvalue

}