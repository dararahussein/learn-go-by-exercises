package api

import (
	"net/http"

	"learn-go-by-exercises/set10/task"
)

type API struct{ Store *task.Store }

func NewHandler(store *task.Store) http.Handler {
	api := &API{Store: store}
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", api.handleTasks)
	mux.HandleFunc("/tasks/", api.handleTaskAction)
	return mux
}

func (a *API) handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO: JSON encode Pending with Content-Type application/json
	case http.MethodPost:
		// TODO: decode {"text":"..."}, reject invalid/empty input, Add, return 201 JSON
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (a *API) handleTaskAction(w http.ResponseWriter, r *http.Request) {
	// TODO: parse /tasks/{id}/done, call Complete, map ErrNotFound to 404
	w.WriteHeader(http.StatusNotImplemented)
}
