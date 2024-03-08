package main

import "fmt"

func main() {
	poem := getRandomPoem()
	fmt.Println(poem)
	sendEmail(poem)
}
