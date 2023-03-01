package main

import "fmt"

func main() {

	Hello:=saySomething()
	fmt.Println("Here is the word return from function ",Hello)
}
func saySomething() string {
	return "Hello world"
}