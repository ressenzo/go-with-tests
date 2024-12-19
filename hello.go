package main

import "fmt"

const englishHelloPreffix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPreffix + name
}

func main() {
	fmt.Println(Hello("world"))
}
