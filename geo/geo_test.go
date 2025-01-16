package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange - подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Ошибка получения города", err.Error())
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected.City, got)
	}

	// Assert - проверка результата с expected
}

func TestGetLocationNoCity(t *testing.T) {
	city := "Londonefddsg"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}
}
