package product

import (
	"net/http"
	"tonylee221b/e-commerce/pkg/jsonutil"
)

type ProductHandler struct {
	svc *ProductService
}

func NewHandler(svc *ProductService) *ProductHandler {
	return &ProductHandler{svc}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req Product
	req, err := jsonutil.ReadJSON[Product](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.svc.CreateProduct(r.Context(), req.ID, req.Name, req.Price.amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jsonutil.WriteJSON(w, http.StatusCreated, map[string]string{"status": "created"})
}

func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.svc.ListProducts(r.Context())
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}
	jsonutil.WriteJSON(w, http.StatusOK, list)
}
