package api

import (
	"net/http"
	"time"

	"OnThisDay/internal/events"
)

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	cfg := events.Config{
		Month:    int(time.Now().Month()),
		Day:      time.Now().Day(),
		Language: "en",
		Type:     "all",
	}

	ev, err := events.Events(cfg)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, ev, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
		return
	}
}
