package main

import "fmt"

func main() {
	element := struct {
		name      string
		particles int
	}{
		name:      "sneha",
		particles: 12,
	}
	fmt.Println(element)
}
