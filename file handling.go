package main

import (
	"fmt"
	"io"
	"os"
)

func readFile (filename string){
	data, err := os.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(data))
}

func main(){
	myfile, err := os.Create("file.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(myfile.Name())
	
	len, err := io.WriteString(myfile, "Hello World")

	if err != nil {
		panic(err)
	}

	fmt.Println(len)
	readFile(myfile.Name())

	defer myfile.Close()
}