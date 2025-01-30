package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	/*	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}*/

	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		_, err := w.Write([]byte("Reset is only allowed in dev environment."))
		if err != nil {
			return
		}
		return
	}

	cfg.fileserverHits.Store(0)
	err := cfg.db.Reset(r.Context())
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Hits reset to 0 and database reset to initial state."))
	if err != nil {
		return
	}
}
