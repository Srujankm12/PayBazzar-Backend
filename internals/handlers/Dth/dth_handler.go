package dth_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	dth_domain "github.com/paybazar-backend/internals/domain/Dth"
	dth_service "github.com/paybazar-backend/internals/service/Dth"
)

type DTHRechargeHandler struct {
	Service *dth_service.DTHRechargeService
}

func NewDTHRechargeHandler(s *dth_service.DTHRechargeService) *DTHRechargeHandler {
	return &DTHRechargeHandler{Service: s}
}

func (h *DTHRechargeHandler) RechargeDTH(c echo.Context) error {
	var req dth_domain.DTHRechargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.ProcessDTHRecharge(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h *DTHRechargeHandler) GetDTHRecharge(c echo.Context) error {
	req := dth_domain.DTHRechargeRequest{
		CustomerID:       c.QueryParam("customer_id"),
		OperatorCode:     parseInt(c.QueryParam("operator_code")),
		Amount:           parseFloat(c.QueryParam("amount")),
		PartnerRequestID: c.QueryParam("partner_request_id"),
		UserVar1:         c.QueryParam("user_var1"),
		UserVar2:         c.QueryParam("user_var2"),
		UserVar3:         c.QueryParam("user_var3"),
	}

	res, err := h.Service.FetchDTHRecharge(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func parseInt(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

func parseFloat(val string) float64 {
	f, _ := strconv.ParseFloat(val, 64)
	return f
}
