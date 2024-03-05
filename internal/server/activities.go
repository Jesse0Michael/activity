package server

import (
	"log/slog"
	"net/http"

	"github.com/jesse0michael/activity/internal/service"
)

func HandleActivities(handler ActivityHandler) http.HandlerFunc {
	type ActivitiesResponse struct {
		Activities []service.Activity `json:"activities"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req service.ActivitiesRequest
		err := decodeRequest(r, &req)
		if err != nil {
			encodeError(w, http.StatusBadRequest, err)
			return
		}

		activities, err := handler.Activities(ctx, req)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get activities", err)
			encodeError(w, http.StatusInternalServerError, err)
			return
		}

		_ = encode(w, http.StatusOK, ActivitiesResponse{Activities: activities})
	}
}

func HandleActivity(handler ActivityHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		activityID := r.PathValue("activityID")

		activity, err := handler.Activity(ctx, activityID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get activity", err)
			encodeError(w, http.StatusInternalServerError, err)
			return
		}

		_ = encode(w, http.StatusOK, *activity)
	}
}

func HandleCreateActivity(handler ActivityHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var req service.Activity
		err := decode(r, &req)
		if err != nil {
			encodeError(w, http.StatusBadRequest, err)
			return
		}

		err = handler.CreateActivity(ctx, req)
		if err != nil {
			slog.ErrorContext(ctx, "failed to create activity", err)
			encodeError(w, http.StatusInternalServerError, err)
			return
		}

		_ = encode(w, http.StatusOK, req)
	}
}

func HandleUpdateActivity(handler ActivityHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		activityID := r.PathValue("activityID")
		var req service.Activity
		err := decode(r, &req)
		if err != nil {
			encodeError(w, http.StatusBadRequest, err)
			return
		}
		req.ID = activityID

		err = handler.UpdateActivity(ctx, activityID, req)
		if err != nil {
			slog.ErrorContext(ctx, "failed to update activity", err)
			encodeError(w, http.StatusInternalServerError, err)
			return
		}

		_ = encode(w, http.StatusOK, req)
	}
}

func HandlePatchActivity(handler ActivityHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		activityID := r.PathValue("activityID")
		var req service.Activity
		err := decode(r, &req)
		if err != nil {
			encodeError(w, http.StatusBadRequest, err)
			return
		}
		req.ID = activityID

		activity, err := handler.PatchActivity(ctx, activityID, req)
		if err != nil {
			slog.ErrorContext(ctx, "failed to patch activity", err)
			encodeError(w, http.StatusInternalServerError, err)
			return
		}

		_ = encode(w, http.StatusOK, activity)
	}
}

func HandleDeleteActivity(handler ActivityHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		activityID := r.PathValue("activityID")

		ok, err := handler.DeleteActivity(ctx, activityID)
		if err != nil {
			slog.ErrorContext(ctx, "failed to delete activity", err)
			encodeError(w, http.StatusInternalServerError, err)
			return
		}

		if ok {
			w.WriteHeader(http.StatusNoContent)
		} else {
			w.WriteHeader(http.StatusGone)
		}

	}
}
