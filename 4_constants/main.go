package main

import "fmt"

func main() {
	const name = "golang"
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(name)
	fmt.Println(port, host)
}