package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/odd1n3rd/default_web/internal/entity"
	"github.com/odd1n3rd/default_web/internal/service"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o entity.Order
	if err := render.DecodeJSON(r.Body, o); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	if err := h.service.CreateOrder(&o); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, o)

}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []entity.Order
	orders, err := h.service.GetAllOrders()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.JSON(w, r, orders)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": err.Error()})
	}

	render.JSON(w, r, order)
}
