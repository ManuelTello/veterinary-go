package main

import "github.com/ManuelTello/veterinary/internal/application"

func main() {
	app := application.New()
	app.StartServer()
}
