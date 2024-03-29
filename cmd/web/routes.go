package main

import (
	"net/http"

	"github.com/exvimmer/lets_go/snippetbox/ui"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(
		func(w http.ResponseWriter, _ *http.Request) {
			app.notFound(w)
		},
	)

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)
	router.HandlerFunc(http.MethodGet, "/ping", ping)

	// unprotected routes
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/about", dynamic.ThenFunc(app.about))
	router.Handler(
		http.MethodGet, "/snippet/view/:id", dynamic.ThenFunc(app.snippetView))
	router.Handler(
		http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignUp))
	router.Handler(
		http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignUpPost))
	router.Handler(http.MethodGet,
		"/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(
		http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	// protected routes
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(
		http.MethodGet, "/snippet/create", protected.ThenFunc(app.snippetCreate))
	router.Handler(
		http.MethodPost,
		"/snippet/create", protected.ThenFunc(app.snippetCreatePost))
	router.Handler(
		http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))
	router.Handler(http.MethodGet, "/account/view",
		protected.ThenFunc(app.accountView))
	router.Handler(http.MethodGet, "/account/password/update",
		protected.ThenFunc(app.accountPasswordUpdate))
	router.Handler(http.MethodPost, "/account/password/update",
		protected.ThenFunc(app.accountPasswordUpdatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
