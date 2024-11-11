package handler

import (
	// "fmt"
	"log"
	"net/http"
)





func (h *AuthHandler) dashboard(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	err := h.tmpl.ExecuteTemplate(w, "dashboard.html", nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Dashboard(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")
	log.Println("Dashboard endpoint accessed token:", token)

	if r.Method == http.MethodGet {
		h.dashboard(w, r)
		return
	}

}
