package electricitybillpayment_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	electricitybillpayment_domain "github.com/paybazar-backend/internals/domain/BillPayments/ElectricityRecharge"
	electricitybillpayment_service "github.com/paybazar-backend/internals/service/BillPayments/ElectricityRecharge"
)

type ElectricityBillPaymentHandler struct {
	Service *electricitybillpayment_service.ElectricityBillPaymentService
}

func NewElectricityBillPaymentHandler(s *electricitybillpayment_service.ElectricityBillPaymentService) *ElectricityBillPaymentHandler {
	return &ElectricityBillPaymentHandler{Service: s}
}

func (h *ElectricityBillPaymentHandler) MakePaymentPOST(c echo.Context) error {
	var req electricitybillpayment_domain.ElectricityBillPaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.MakePaymentPOST(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *ElectricityBillPaymentHandler) MakePaymentGET(c echo.Context) error {
	req := electricitybillpayment_domain.ElectricityBillPaymentRequest{
		P1:               c.QueryParam("p1"),
		P2:               c.QueryParam("p2"),
		P3:               c.QueryParam("p3"),
		CustomerEmail:    c.QueryParam("customer_email"),
		PartnerRequestID: c.QueryParam("partner_request_id"),
		UserVar1:         c.QueryParam("user_var1"),
		UserVar2:         c.QueryParam("user_var2"),
		UserVar3:         c.QueryParam("user_var3"),
	}

	req.OperatorCode, _ = strconv.Atoi(c.QueryParam("operator_code"))
	req.Amount, _ = strconv.ParseFloat(c.QueryParam("amount"), 64)

	res, err := h.Service.MakePaymentGET(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
