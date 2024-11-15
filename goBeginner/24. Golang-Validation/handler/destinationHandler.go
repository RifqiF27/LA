package handler

import (
	"net/http"
	"strconv"
	"travelika/service"
	"travelika/utils"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type DestinationHandler struct {
	Service *service.DestinationService
	Log     *zap.Logger
}

func NewDestinationHandler(service *service.DestinationService, logger *zap.Logger) *DestinationHandler {
	return &DestinationHandler{Service: service, Log: logger}
}

func (h *DestinationHandler) GetDestination(w http.ResponseWriter, r *http.Request) {

	h.Log.Info("Handler: Received request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	date := r.URL.Query().Get("date")
	eventName := r.URL.Query().Get("event_name")
	location := r.URL.Query().Get("location")
	orderBy := r.URL.Query().Get("low_to_high")
	orderAsc := r.URL.Query().Get("orderDesc") == "true"
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if limit == 0 {
		limit = 6
	}
	if page == 0 {
		page = 1
	}

	destinations, totalItems, totalPages, err := h.Service.GetAllEvents(eventName, location, date, orderBy, orderAsc, limit, page)
	if err != nil {
		h.Log.Error("Handler: Error getting events", zap.Error(err))
		utils.SendJSONResponse(w, false, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(destinations) == 0 {
		h.Log.Warn("Handler: No destinations found", zap.String("date", date), zap.String("eventName", eventName), zap.String("location", location))
		utils.SendJSONResponse(w, false, http.StatusNotFound, "No destinations found", nil)
		return
	}

	utils.SendJSONResponsePagination(w, true, page, limit, totalItems, totalPages, http.StatusOK, "", destinations)
}

func (h *DestinationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	h.Log.Info("Handler: Received request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	item, err := h.Service.GetById(id)
	if err != nil {
		h.Log.Error("Handler: Destination not found", zap.Error(err) )
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Destination not found", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", item)
}
func (h *DestinationHandler) GetTourPlansByEventID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	h.Log.Info("Handler: Received request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	item, err := h.Service.GetTourPlansByEventID(id)
	if err != nil {
		h.Log.Error("Handler: event not found", zap.Error(err) )
		utils.SendJSONResponse(w, false, http.StatusNotFound, "event not found", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", item)
}
func (h *DestinationHandler) GetLocationsByDestinationID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	h.Log.Info("Handler: Received request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	item, err := h.Service.GetLocationsByDestinationID(id)
	if err != nil {
		h.Log.Error("Handler: destination not found", zap.Error(err) )
		utils.SendJSONResponse(w, false, http.StatusNotFound, "Destination not found", nil)
		return
	}

	utils.SendJSONResponse(w, true, http.StatusOK, "", item)
}
