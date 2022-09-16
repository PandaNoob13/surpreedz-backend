package main

import (
	"surpreedz-backend/delivery"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config.env")
	if err != nil {
		panic(err)
	}
	delivery.Server().Run()

}
