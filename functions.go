package main

import "fmt"

// func isEven(num int) string{
// 	if num % 2 == 0{
// 		return "Even"
// 	}else{
// 		return "Odd"
// 	}
// }

// func arrSum(arr []int) int {
// 	var sum int
// 	for _, val := range arr {
// 		sum += val
// 	}
// 	return sum
// }

func arithmaticOperations(a int, b int) (int, int, int, float32){
	var add = a+b
	var sub = a-b
	var mul = a*b
	var div = float32(a)/float32(b)

	return add, sub, mul, div
}

func main() {
	// fmt.Printf("%d is %s\n", 5, isEven(5))

	// arr := []int{2, 4, 6, 8}

	// fmt.Println(arrSum(arr))

	// add,sub,mul,div := arithmaticOperations(5, 8)

	// fmt.Println(add,sub,mul,div)
}
