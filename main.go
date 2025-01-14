package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Weather API")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Printf("Город, %s!\n", *city)
	fmt.Println(*format)

	r := strings.NewReader("Привет! я поток данных!")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}
}
