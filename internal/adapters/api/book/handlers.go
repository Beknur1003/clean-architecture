package book

// handlers - controllers
// Здесь может быть REST API,Graph QL, RP и что угодно
import (
	"net/http"

	"github.com/Beknur1003/clean-architecture.git/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "/book/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)

	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
