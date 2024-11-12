package collections

type Book struct {
	ID       int      `json:"id"`
	BookCode string   `json:"book_code"`
	Title    string   `json:"title"`
	Category Category `json:"category_id"`
	Author   string   `json:"author"`
	Price    float64  `json:"price"`
	Discount float64  `json:"discount"`
	Cover    string   `json:"cover"`
	FileBook string   `json:"file_book"`
}
