package main

import "fmt"

func main() {
	db := NewDb()

	s := NewServer(db)

	fmt.Println("listening on port 8080")
	s.ListenAndServe()
}
