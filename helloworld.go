package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func sayHello() string {
	log.Info("saying hello...")
	return "Hello, world"
}

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println(sayHello())
}