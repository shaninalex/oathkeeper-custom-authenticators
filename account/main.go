package main

import (
	"os"
	"strconv"
)

var (
	PORT = os.Getenv("PORT")
)

func main() {
	port, err := strconv.Atoi(PORT)
	if err != nil {
		panic(err)
	}

	app, err := InitApplication()
	if err != nil {
		panic(err)
	}
	app.Run(port)
}
