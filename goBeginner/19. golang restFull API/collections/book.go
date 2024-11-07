package collections

type Book struct {
    ID       int     `json:"id"`
    BookCode string  `json:"book_code"`
    Name     string  `json:"name"`
    Type     string  `json:"type"`
    Author   string  `json:"author"`
    Price    float64 `json:"price"`
    Discount float64 `json:"discount"`
    Cover    string  `json:"cover"`
}