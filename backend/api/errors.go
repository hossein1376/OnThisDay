package api

import (
	"log/slog"
	"net/http"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	err := writeJSON(w, status, message, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error("Server Error",
		slog.String("path:", r.URL.Path),
		slog.String("method:", r.Method),
		slog.String("error:", err.Error()),
	)
	message := "Internal Server Error"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "Not Found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "Method Not Allowed"
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, message interface{}) {
	app.errorResponse(w, r, http.StatusBadRequest, message)
}

func (app *application) unauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	message := "Unauthorized"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	message := "Forbidden"
	app.errorResponse(w, r, http.StatusForbidden, message)
}
