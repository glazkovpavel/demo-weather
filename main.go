package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	fmt.Println("Welcome to the Weather API")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Printf("Город, %s!\n", *city)
	fmt.Println(*format)
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(geoData)

	weatherResp := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherResp)
}

//r := strings.NewReader("Привет! я поток данных!")
//block := make([]byte, 4)
//for {
//	_, err := r.Read(block)
//	fmt.Printf("%q\n", block)
//	if err == io.EOF {
//		break
//	}
//}
