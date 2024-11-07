package router

// import (
// 	"book-store/database"

// 	"github.com/go-chi/chi/v5"
// )

// func NewRouter() *chi.Mux {
// 	r := chi.NewRouter()
// 	db := database.NewPostgresDB()

// 	repo := repository.NewProductRepository(db)
// 	srv := service.NewProductService(repo)
// 	h := handler.NewProductHandler(srv)

// 	r.Use(middleware.Logger)
// 	r.Use(middleware.BasicAuth)

// 	r.Get("/products", h.GetAllProducts)
// 	r.Get("/products/{id}", h.GetProductByID)
// 	r.Post("/products", h.CreateProduct)
// 	r.Put("/products/{id}", h.UpdateProduct)
// 	r.Delete("/products/{id}", h.DeleteProduct)

// 	return r
// }
