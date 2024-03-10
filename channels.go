package main
import "fmt"

func sum(a, b int, ch chan int) {
	c := a+b
	ch <- c
}

func main(){
	ch := make(chan int)
	go sum(40, 80, ch)
	output := <-ch
	fmt.Println(output)
}