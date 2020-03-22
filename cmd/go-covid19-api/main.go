package main

import "github.com/kosegor/go-covid19-api/app/server"

func main() {
	router := server.CreateServer()
	router.Run()
}
