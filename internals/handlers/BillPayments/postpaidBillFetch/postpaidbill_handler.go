package postpaidbill_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	postpaidbill_domain "github.com/paybazar-backend/internals/domain/BillPayments/postpaidBillFetch"
	postpaidbill_service "github.com/paybazar-backend/internals/service/BillPayments/postpaidBillFetch"
)

type PostpaidBillFetchHandler struct {
	Service *postpaidbill_service.PostpaidBillFetchService
}

func NewPostpaidBillFetchHandler(s *postpaidbill_service.PostpaidBillFetchService) *PostpaidBillFetchHandler {
	return &PostpaidBillFetchHandler{Service: s}
}

func (h *PostpaidBillFetchHandler) SavePostPaidBill(c echo.Context) error {
	var req postpaidbill_domain.PostpaidBillRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.SavePostPaidBill(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h *PostpaidBillFetchHandler) FetchBill(c echo.Context) error {
	mobileNo := c.QueryParam("mobile_no")
	operatorCodeStr := c.QueryParam("operator_code")

	if mobileNo == "" || operatorCodeStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing query params"})
	}

	operatorCode, err := strconv.Atoi(operatorCodeStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid operator_code"})
	}

	req := &postpaidbill_domain.PostpaidBillRequest{
		OperatorCode: operatorCode,
		MobileNo:     mobileNo,
	}

	res, err := h.Service.FetchBill(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
