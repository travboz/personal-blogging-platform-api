package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

// The logError() method is a generic helper for logging an error message along
// with the current request method and URL as attributes in the log entry.
func logError(logger *slog.Logger, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	logger.Error(err.Error(), "method", method, "uri", uri)
}

// The errorResponse() method is a generic helper for sending JSON-formatted error
// messages to the client with a given status code.
func errorResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	// Write the response using the writeJSON() helper. If this happens to return an

	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := writeJSON(w, status, env, nil)
	if err != nil {
		logError(logger, r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func serverErrorResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request, err error) {
	logError(logger, r, err)
	message := "the server encountered a problem and could not process your request"
	errorResponse(logger, w, r, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send a 404 Not Found status code and
// JSON response to the client.
func notFoundResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(logger, w, r, http.StatusNotFound, message)
}

// The methodNotAllowedResponse() method will be used to send a 405 Method Not Allowed
// status code and JSON response to the client.
func methodNotAllowedResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(logger, w, r, http.StatusMethodNotAllowed, message)
}

func badRequestResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(logger, w, r, http.StatusBadRequest, err.Error())
}

func failedValidationResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request, errors map[string]string) {
	errorResponse(logger, w, r, http.StatusUnprocessableEntity, errors)
}

func editConflictResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	errorResponse(logger, w, r, http.StatusConflict, message)
}

func rateLimitExceededResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "rate limit exceeded"
	errorResponse(logger, w, r, http.StatusTooManyRequests, message)
}

func invalidCredentialsResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	errorResponse(logger, w, r, http.StatusUnauthorized, message)
}

func invalidAuthenticationTokenResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	/*
		We’re including a WWW-Authenticate: Bearer header here to help inform or
		remind the client that we expect them to authenticate using a bearer token.
	*/
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	errorResponse(logger, w, r, http.StatusUnauthorized, message)
}

func authenticationRequiredResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	errorResponse(logger, w, r, http.StatusUnauthorized, message)
}

func inactiveAccountResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	errorResponse(logger, w, r, http.StatusForbidden, message)
}

func notPermittedResponse(logger *slog.Logger, w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	errorResponse(logger, w, r, http.StatusForbidden, message)
}
