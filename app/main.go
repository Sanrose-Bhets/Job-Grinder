package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Hello, World!")

	port := os.Getenv("PORT")

	fmt.Println(port)
}