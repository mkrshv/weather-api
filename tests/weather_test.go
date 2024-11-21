package tests

import (
	"fmt"
	"testing"
	"weatherapi"
)

func TestGetWeather(t *testing.T) {
	client := weatherapi.NewClient(API_KEY, BASE_URL)
	res, err := client.GetWeather("Moscow")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.Format())
}
