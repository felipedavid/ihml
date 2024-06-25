package handlers

import "net/http"

func DefineRoutes() http.Handler {
	mux := http.NewServeMux()

	return mux
}
