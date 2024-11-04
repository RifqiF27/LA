package model

type PaginationRequest struct {
    Page  int `json:"page"`
    Limit int `json:"limit"`
    FilterStock bool `json:"filter_stock,omitempty"`
    SearchName string `json:"search_name,omitempty"`

}