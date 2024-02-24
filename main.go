package main

import "fmt"

func main() {
	server := NewAPIServer(":8080")
	server.Run()
	fmt.Println("let's fucking GO!")
}
