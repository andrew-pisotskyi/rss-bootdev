package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type OK struct {
		Status string `json:"status"`
	}

	respondWithJSON(w, http.StatusOK, OK{"ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
