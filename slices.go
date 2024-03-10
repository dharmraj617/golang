package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6,7,8)
	fmt.Println(slice)
}
