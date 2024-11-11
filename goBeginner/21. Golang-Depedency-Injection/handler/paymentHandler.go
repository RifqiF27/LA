package handler

import (
	"book-store/collections"
	"book-store/service"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type PaymentMethodHandler struct {
	service service.PaymentMethodService
}

func NewPaymentMethodHandler(service service.PaymentMethodService) *PaymentMethodHandler {
	return &PaymentMethodHandler{service}
}

func (h *PaymentMethodHandler) CreatePaymentMethod(w http.ResponseWriter, r *http.Request) {
	domain := strings.Join([]string{"http://", r.Host},"") 
   if err := r.ParseMultipartForm(10 << 20); err != nil {
	   http.Error(w, "File too large or invalid form data", http.StatusBadRequest)
	   return
   }

   name := r.FormValue("name")
   isActiveStr := r.FormValue("is_active")
   isActive, err := strconv.ParseBool(isActiveStr)
   if err != nil {
	   http.Error(w, "Invalid value for is_active", http.StatusBadRequest)
	   return
   }

   file, handler, err := r.FormFile("photo")
   if err != nil {
	   http.Error(w, "Failed to read photo file", http.StatusBadRequest)
	   return
   }
   defer file.Close()

   filePath := filepath.Join("static", handler.Filename)
   dst, err := os.Create(filePath)
   if err != nil {
	   http.Error(w, "Failed to save photo", http.StatusInternalServerError)
	   return
   }
   defer dst.Close()

   if _, err := io.Copy(dst, file); err != nil {
	   http.Error(w, "Failed to save photo", http.StatusInternalServerError)
	   return
   }

   method := &collections.PaymentMethod{
	   Name:     name,
	   PhotoURL: strings.Join([]string{domain, "/static/", handler.Filename},"") ,
	   IsActive: isActive,
   }

   if err := h.service.CreatePaymentMethod(r.Context(), method); err != nil {
	   http.Error(w, err.Error(), http.StatusInternalServerError)
	   return
   }

   w.WriteHeader(http.StatusCreated)
   json.NewEncoder(w).Encode(method)
}

func (h *PaymentMethodHandler) ListPaymentMethods(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentMethodHandler) GetPaymentMethod(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentMethodHandler) UpdatePaymentMethod(w http.ResponseWriter, r *http.Request) {

}

func (h *PaymentMethodHandler) DeletePaymentMethod(w http.ResponseWriter, r *http.Request) {

}
