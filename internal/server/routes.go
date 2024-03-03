package server

import (
	"net/http"
)

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /activities", nil)
	mux.HandleFunc("POST /activities", nil)
	mux.HandleFunc("GET /activivites/{activityID}", nil)
	mux.HandleFunc("PUT /activivites/{activityID}", nil)
	mux.HandleFunc("PATCH /activivites/{activityID}", nil)
	mux.HandleFunc("DELETE /activivites/{activityID}", nil)

	return mux
}
