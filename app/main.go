package main

import (
	"fmt"
	"os"
)

func main(){

	//loading all environmental variables
	godotenv.Load()

	//loading the dabase connection
	DBURL=os.Getenv("DB_URL")

	fmt.Println("Hello, World!")

	port := os.Getenv("PORT")

	fmt.Println(port)
}