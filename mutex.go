package main

import (
	"fmt"
	"sync"
)

func incr(x *int, mu *sync.Mutex, wg *sync.WaitGroup){
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	*x++
}

func decr(x *int, mu *sync.Mutex, wg *sync.WaitGroup){
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	*x--
}

func main(){
	x:=8
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go incr(&x, &mu, &wg)
	wg.Add(1)
	go decr(&x, &mu, &wg)
	fmt.Println(x)
	wg.Done()
}