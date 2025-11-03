package payouthandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	payoutdomain "github.com/paybazar-backend/internals/domain/payout"
	payoutservice "github.com/paybazar-backend/internals/service/payout"
)

type PayoutHandler struct {
	Service *payoutservice.PayoutService
}

func NewPayoutHandler(service *payoutservice.PayoutService) *PayoutHandler {
	return &PayoutHandler{Service: service}
}
func (h *PayoutHandler) InitiatePayout(c echo.Context) error {
	var req payoutdomain.PayoutRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	res, err := h.Service.InitiatePayout(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
