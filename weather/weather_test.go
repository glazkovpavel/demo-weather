package weather_test

import (
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

func TestGetWeather(t *testing.T) {
	expected := "Moscow"
	groData := geo.GeoData{
		City: expected,
	}
	format := 3
	result, err := weather.GetWeather(groData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получение %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{
		name:   "Big format",
		format: 145,
	},
	{
		name:   "0 format",
		format: 0,
	},
	{
		name:   "Minus format",
		format: -1,
	},
}

func TestGetWeatherWithInvalidFormat(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			expected := "Moscow"
			groData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(groData, testCase.format)
			if err != weather.ErrInvalidFormat {
				t.Errorf("Ожидалось %v, получение %v", weather.ErrInvalidFormat, err)
			}
		})
	}

}
