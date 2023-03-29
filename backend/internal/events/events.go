package events

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// OnThisDay is a struct type that matches the JSON structure of the API response
type OnThisDay struct {
	Selected []Event `json:"selected"`
	Events   []Event `json:"events"`
	Holidays []Event `json:"holidays"`
	Births   []Event `json:"births"`
	Deaths   []Event `json:"deaths"`
}

// Event is a struct type that represents an event
type Event struct {
	Text  string `json:"text"`
	Pages []Page `json:"pages"`
	Year  int    `json:"year"`
}

// Page is a struct type that represents a page
type Page struct {
	Title       string     `json:"title"`
	Pageid      int        `json:"pageid"`
	Thumbnail   *Image     `json:"thumbnail,omitempty"`
	Original    *Image     `json:"originalimage,omitempty"`
	Timestamp   string     `json:"timestamp,omitempty"`   // Not present for holidays
	Description string     `json:"description,omitempty"` // Not present for holidays
	ContentURL  ContentURL `json:"content_urls"`
	Extract     string     `json:"extract"`
}

// ContentURL is a struct type that represents content URLs
type ContentURL struct {
	Desktop Content `json:"desktop"`
	Mobile  Content `json:"mobile"`
}

// Content is a struct type that represents content
type Content struct {
	Page string `json:"page"`
}

// Image is a struct type that represents an image
type Image struct {
	Source string `json:"source"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Config is a struct type that represents the configurations
type Config struct {
	Month    int
	Day      int
	Language string
	Type     string
}

func Events(cfg Config) (*OnThisDay, error) {
	url := fmt.Sprintf("https://api.wikimedia.org/feed/v1/wikipedia/%s/onthisday/%s/%v/%v", cfg.Language, cfg.Type, cfg.Month, cfg.Day)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var otd OnThisDay

	err = json.NewDecoder(resp.Body).Decode(&otd)
	if err != nil {
		return nil, err
	}

	return &otd, nil
}
