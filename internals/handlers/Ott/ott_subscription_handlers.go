package ott_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	ott_domain "github.com/paybazar-backend/internals/domain/Ott"
	ott_service "github.com/paybazar-backend/internals/service/Ott"
)

type OTTSubscriptionHandler struct {
	Service *ott_service.OTTSubscriptionService
}

func NewOTTSubscriptionHandler(s *ott_service.OTTSubscriptionService) *OTTSubscriptionHandler {
	return &OTTSubscriptionHandler{Service: s}
}
func (h *OTTSubscriptionHandler) CreateSubscription(c echo.Context) error {
	var req ott_domain.OTTSubscriptionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.CreateSubscription(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h *OTTSubscriptionHandler) CreateSubscriptionGET(c echo.Context) error {
	req := ott_domain.OTTSubscriptionRequest{
		MobileNo:         c.QueryParam("mobile_no"),
		CustomerEmail:    c.QueryParam("customer_email"),
		PartnerRequestID: c.QueryParam("partner_request_id"),
		UserVar1:         c.QueryParam("user_var1"),
		UserVar2:         c.QueryParam("user_var2"),
		UserVar3:         c.QueryParam("user_var3"),
	}

	operatorCode, _ := strconv.Atoi(c.QueryParam("operator_code"))
	amount, _ := strconv.ParseFloat(c.QueryParam("amount"), 64)
	req.OperatorCode = operatorCode
	req.Amount = amount

	res, err := h.Service.CreateSubscriptionGET(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
