package handler

import (
	"cinema/model"
	"cinema/service"
	"cinema/utils"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CinemaHandler struct {
	Service service.CinemaService
}

func NewCinemaHandler(service service.CinemaService) *CinemaHandler {
	return &CinemaHandler{Service: service}
}

func (h *CinemaHandler) GetAllCinemas(w http.ResponseWriter, r *http.Request) {
	cinemas, err := h.Service.GetAllCinemas()
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendJSONResponse(w, http.StatusOK, "", cinemas)
}

func (h *CinemaHandler) GetCinemaByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "cinemaId"))
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	cinema, err := h.Service.GetCinemaByID(int(id))
	if err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendJSONResponse(w, http.StatusOK, "", cinema)

}

func (h *CinemaHandler) GetSeats(w http.ResponseWriter, r *http.Request) {
	cinemaID, err := strconv.Atoi(chi.URLParam(r, "cinemaId"))
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "cinema ID not found", nil)
		return
	}

	date := r.URL.Query().Get("date")
	time := r.URL.Query().Get("time")

	seats, err := h.Service.GetAvailableSeats(int(cinemaID), date, time)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, "No available seats found for the specified cinema and schedule.", nil)
		return
	}

	// Convert to desired response format
	var seatResponses []model.SeatResponse
	for _, seat := range seats {
		seatResponses = append(seatResponses, model.SeatResponse{
			SeatID: strconv.Itoa(seat.ID), // or format as desired
			Status: seat.Status,
		})
	}

	utils.SendJSONResponse(w, http.StatusOK, "", seatResponses)
}

// func (h *CinemaHandler) GetSeats(w http.ResponseWriter, r *http.Request) {
// 	cinemaID, err := strconv.Atoi(chi.URLParam(r, "cinemaId"))
// 	if err != nil {
// 		utils.SendJSONResponse(w, http.StatusBadRequest, "cinema ID not found", nil)
// 		return
// 	}

// 	date := r.URL.Query().Get("date")
// 	time := r.URL.Query().Get("time")

// 	seats, err := h.Service.GetAvailableSeats(int(cinemaID), date, time)
// 	if err != nil {
// 		utils.SendJSONResponse(w, http.StatusNotFound, "No available seats found for the specified cinema and schedule.", nil)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	utils.SendJSONResponse(w, http.StatusOK, "", seats)

// }
