package model

type PaginationRequest struct {
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
	SearchThread string `json:"search_thread,omitempty"`
}
