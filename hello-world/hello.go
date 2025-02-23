package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("World"))
}

func Hello(name string) string {
	return englishHelloPrefix + name
}
