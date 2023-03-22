package api

import (
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	err := writeJSON(w, status, message, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s %s %s", r.Method, r.URL.Path, err)
	message := "Internal Server Error"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "Not Found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := "Method Not Allowed"
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "Bad Request"
	errorResponse(w, r, http.StatusBadRequest, message)
}

func unauthorizedResponse(w http.ResponseWriter, r *http.Request) {
	message := "Unauthorized"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

func forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	message := "Forbidden"
	errorResponse(w, r, http.StatusForbidden, message)
}
