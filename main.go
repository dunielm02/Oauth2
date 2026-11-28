package main

import "fmt"

func main() {
	NewDb()

	s := NewServer()

	fmt.Println("listening on port 8080")
	s.ListenAndServe()
}
