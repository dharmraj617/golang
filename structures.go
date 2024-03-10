package main

import "fmt"

type Player struct{
	first_name string
	last_name string
	country string
	age int
}

func main(){
	// var player1 = Player{"Virat", "Kohli", "India", 30}
	var player2 = Player{first_name: "Dharmraj", country: "India", last_name: "Patil"}
	player2.first_name = "raj"
	fmt.Println(player2)
}