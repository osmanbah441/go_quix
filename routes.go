package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// todo: add your handlers
	// router.HandlerFunc(http.MethodGet, "/path", app.handler)

	return app.recoverPanic(app.rateLimit(router))
}
