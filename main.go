package main

import "fmt"

func main() {
	var str string
	fmt.Println("Tiny Menu")
	for {
		fmt.Scan(&str)
		switch str {
		case "version":
			fmt.Println("Version 1.0")
		case "quit":
			fmt.Println("bye")
			return
		}
	}
}
