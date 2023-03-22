package events

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

// OnThisDay is a struct type that matches the JSON structure of the API response
type OnThisDay struct {
	Selected []Event `json:"selected"`
	Births   []Event `json:"births"`
	Deaths   []Event `json:"deaths"`
	Events   []Event `json:"events"`
	Holidays []Event `json:"holidays"`
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

func Events() (string, error) {
	var cfg Config
	flag.StringVar(&cfg.Language, "lang", "en", "Language code")
	flag.StringVar(&cfg.Type, "type", "all", "Type of events; all, selected, births, deaths, events, holidays")
	flag.IntVar(&cfg.Month, "month", int(time.Now().Month()), "Month")
	flag.IntVar(&cfg.Day, "day", time.Now().Day(), "Day")
	flag.Parse()

	otd, err := getEvents(cfg)
	if err != nil {
		return "", err
	}

	//fmt.Println("Selected events:")
	//for _, e := range otd.Selected {
	//	fmt.Printf("%d: %s\n", e.Year, e.Text)
	//}

	file, err := os.Create("response.html")
	if err != nil {
		return "", err
	}
	defer file.Close()

	html := "<html><head><title>On This Day</title></head><body><h2>Selected events</h2><ul>"
	for _, e := range otd.Selected {
		html += fmt.Sprintf("<li>%d: %s</li>", e.Year, e.Text)
	}
	html += "</ul><h2>Births</h2><ul>"
	for _, e := range otd.Births {
		html += fmt.Sprintf("<li>%d: %s</li>", e.Year, e.Text)
	}
	html += "</ul><h2>Deaths</h2><ul>"
	for _, e := range otd.Deaths {
		html += fmt.Sprintf("<li>%d: %s</li>", e.Year, e.Text)
	}
	html += "</ul><h2>Events</h2><ul>"
	for _, e := range otd.Events {
		html += fmt.Sprintf("<li>%d: %s</li>", e.Year, e.Text)
	}
	html += "</ul><h2>Holidays</h2><ul>"
	for _, e := range otd.Holidays {
		html += fmt.Sprintf("<li>%s</li>", e.Text)
	}
	html += "</ul></body></html>"

	_, err = file.WriteString(html)
	if err != nil {
		return "", err
	}

	return html, nil
}

func getEvents(cfg Config) (*OnThisDay, error) {
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
