package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"OnThisDay/internal/events"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

func (app *application) GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	cfg := events.Config{
		Month:    int(time.Now().Month()),
		Day:      time.Now().Day(),
		Language: "en",
		Type:     "all",
	}

	// check if data has already been cashed inside Redis
	result, err := RDB.Get(context.Background(), fmt.Sprintf("month%vday%v", cfg.Month, cfg.Day)).Result()
	if err != redis.Nil && result != "" {
		err = writeStringJSON(w, http.StatusOK, result, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	ev, err := events.Events(cfg)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	js, err := json.Marshal(ev)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// cashing data in Redis for 24 hours
	err = RDB.SetNX(context.Background(), fmt.Sprintf("month%vday%v", cfg.Month, cfg.Day), string(js), 24*time.Hour).Err()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, ev, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) GetLanguageEventHandler(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	if lang == "" {
		app.badRequestResponse(w, r, "language is required")
		return
	}
	if len(lang) != 2 {
		app.badRequestResponse(w, r, "language must be 2 characters long")
		return
	}

	cfg := events.Config{
		Month:    int(time.Now().Month()),
		Day:      time.Now().Day(),
		Language: lang,
		Type:     "all",
	}

	// check if data has already been cashed inside Redis
	result, err := RDB.Get(context.Background(), fmt.Sprintf("month%vday%vlang%v", cfg.Month, cfg.Day, cfg.Language)).Result()
	if err != redis.Nil && result != "" {
		err = writeStringJSON(w, http.StatusOK, result, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	ev, err := events.Events(cfg)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	js, err := json.Marshal(ev)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// cashing data in Redis for 24 hours
	err = RDB.SetNX(context.Background(), fmt.Sprintf("month%vday%vlang%v", cfg.Month, cfg.Day, cfg.Language), string(js), 24*time.Hour).Err()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, ev, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
