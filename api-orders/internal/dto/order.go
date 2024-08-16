package dto

import (
	enum "api-orders/internal/enum"
)

type OrderRequest struct {
	// State  enum.EOrderState `form:"state" binding:"required"`
	// UserId string `form:"userId" binding:"required"`
}

type OrderResponse struct {
	Id     string           `json:"_id"`
	State  enum.EOrderState `json:"state"`
	UserId string           `json:"userId"`
}
