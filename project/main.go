package main

import (
	"pahamify/project/controller"
	"pahamify/project/repository"
	"pahamify/project/service"
)

func main() {
	controller.Init(service.Init(repository.Init()))
}
