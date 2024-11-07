package collections

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"-"`
	Data    interface{} `json:"data,omitempty"`
}
