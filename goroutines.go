package main

import (
	"fmt"
	"sync"
	"time"
)

func printNum(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0; i <= 10; i++{
		fmt.Print(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printChar(wg *sync.WaitGroup){
	defer wg.Done()
	for i := 'a'; i < 'e'; i++{
		fmt.Printf("%c", i)
		time.Sleep(100 * time.Millisecond)

	}
}

func main(){
	var wg sync.WaitGroup

	wg.Add(1)
	go printNum(&wg)
	wg.Add(1)
	go printChar(&wg)
	
	wg.Wait()

}