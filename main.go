package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Weather API")
	name := flag.String("name", "Имя по умолчанию", "Имя пользователя")

	flag.Parse()

	fmt.Printf("Hello, %s!\n", *name)

}
