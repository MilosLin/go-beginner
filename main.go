package main

import (
	"fmt"

	"go-beginner/env"
)

func main() {
	fmt.Println("version:" + env.Version)
	fmt.Println("BuildTime:" + env.BuildTime)
}
