package main

import "fmt"

func main() {

	Hello,earth :=saySomething()
	fmt.Println("Here is the word return from function ",Hello,earth)
}
func saySomething() (string,string) {
	return "Hello world", "This is new in go lang"
}