package main

import (
	"OnThisDay/internal/events"
	"fmt"
	"net/smtp"
)

func main() {
	html, err := events.Events()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(html)
}
