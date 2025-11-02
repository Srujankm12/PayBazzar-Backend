package postpaidmobilerecharge_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	postpaidmobilerecharge_domain "github.com/paybazar-backend/internals/domain/BillPayments/postpaidmobilerecharge"
	postpaidmobilerecharge_service "github.com/paybazar-backend/internals/service/BillPayments/postpaidmobilerecharge"
)

type PostpaidRechargeHandler struct {
	Service *postpaidmobilerecharge_service.PostpaidRechargeService
}

func NewPostpaidRechargeHandler(s *postpaidmobilerecharge_service.PostpaidRechargeService) *PostpaidRechargeHandler {
	return &PostpaidRechargeHandler{Service: s}
}

func (h *PostpaidRechargeHandler) RechargePostpaid(c echo.Context) error {
	var req postpaidmobilerecharge_domain.PostpaidMobileRechargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.ProcessPostpaidRecharge(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *PostpaidRechargeHandler) FetchPostpaidRecharge(c echo.Context) error {
	operatorCode, _ := strconv.Atoi(c.QueryParam("operator_code"))
	amount, _ := strconv.ParseFloat(c.QueryParam("amount"), 64)
	circle, _ := strconv.Atoi(c.QueryParam("circle"))
	rechargeType, _ := strconv.Atoi(c.QueryParam("recharge_type"))

	req := &postpaidmobilerecharge_domain.PostpaidMobileRechargeRequest{
		MobileNo:         c.QueryParam("mobile_no"),
		OperatorCode:     operatorCode,
		Amount:           amount,
		PartnerRequestID: c.QueryParam("partner_request_id"),
		Circle:           circle,
		RechargeType:     rechargeType,
		UserVar1:         c.QueryParam("user_var1"),
		UserVar2:         c.QueryParam("user_var2"),
		UserVar3:         c.QueryParam("user_var3"),
	}

	res, err := h.Service.FetchPostpaidRecharge(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
