package server

import (
	"net/http"
)

func routes(activityService ActivityHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /activities", HandleActivities(activityService))
	mux.HandleFunc("POST /activities", HandleCreateActivity(activityService))
	mux.HandleFunc("GET /activities/{activityID}", HandleActivity(activityService))
	mux.HandleFunc("PUT /activities/{activityID}", HandleUpdateActivity(activityService))
	mux.HandleFunc("PATCH /activities/{activityID}", HandlePatchActivity(activityService))
	mux.HandleFunc("DELETE /activities/{activityID}", HandleDeleteActivity(activityService))

	return mux
}
