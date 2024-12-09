package main

import "fmt"

func main() {
	//arr := [20]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"}

	/*
			for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i])
			}

		for _, v := range arr {
			fmt.Println(v)
		}

		fmt.Println(arr[0:10])
	*/
	oldArr := [10]string{"A", "b", "c", "D", "E", "f", "g", "h", "I", "J"}
	slc := oldArr[0:5]
	slc[0] = "START"
	slc = append(slc, "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T")

	fmt.Println(len(slc), cap(slc), slc)

}
