package main

import service "golang-rest-service-template"

func main() {
	if err := service.Run(); err != nil {
		panic(err)
	}
}
