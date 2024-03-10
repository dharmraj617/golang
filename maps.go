package main

import "fmt"

func main() {
	languages := make(map[int]string)
	languages[0] = "Go"
	languages[1] = "C++"
	languages[4] = "Python"
	fmt.Println(languages[1])
	delete(languages, 1)
	fmt.Println(languages)


}
