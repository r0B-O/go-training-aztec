package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range arr {
		if v%2 == 0 {
			fmt.Println("Even Number")
		} else {
			fmt.Println("Odd Number")
		}
	}
}
