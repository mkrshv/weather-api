package weatherapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type APIResponse struct {
	Current WeatherData `json:"current"`
}

type WeatherData struct {
	TempC      float32 `json:"temp_c"`
	FeelsLikeC float32 `json:"feelslike_c"`
	Condition  struct {
		Text string `json:"text"`
	} `json:"condition"`
}

type Client struct {
	apiKey  string
	baseURL string
}

func NewClient(apiKey, baseURL string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: baseURL,
	}
}

func (c *Client) GetWeather(city string) (*APIResponse, error) {
	var res APIResponse
	buf := new(bytes.Buffer)

	req := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no&lang=ru", c.baseURL, c.apiKey, city)

	resp, err := http.Get(req)
	if err != nil {
		return &res, err
	}

	fmt.Println(resp.Body)

	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return &res, err
	}

	if err = json.Unmarshal(buf.Bytes(), &res); err != nil {
		return &res, err
	}

	return &res, nil
}

func (ar *APIResponse) Format() string {
	return fmt.Sprintf("Сейчас %s, температура воздуха: %.1f°C, ощущается как %.1f°C", strings.ToLower(ar.Current.Condition.Text), ar.Current.TempC, ar.Current.FeelsLikeC)
}
