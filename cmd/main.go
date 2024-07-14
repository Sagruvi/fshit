package main

import "awesomeProject/internal/app"

func main() {
	var a app.App
	err := a.Run()
	if err != nil {
		panic(err)
	}
}
