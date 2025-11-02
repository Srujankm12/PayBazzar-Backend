package electricitybillfetch_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	electricitybillfetch_domain "github.com/paybazar-backend/internals/domain/BillPayments/ElectricityBill"
	electricitybillfetch_service "github.com/paybazar-backend/internals/service/BillPayments/ElectricityBill"
)

type ElectricityBillFetchHandler struct {
	Service *electricitybillfetch_service.ElectricityBillFetchService
}

func NewElectricityBillFetchHandler(s *electricitybillfetch_service.ElectricityBillFetchService) *ElectricityBillFetchHandler {
	return &ElectricityBillFetchHandler{Service: s}
}
func (h *ElectricityBillFetchHandler) SaveElectricityBill(c echo.Context) error {
	var req electricitybillfetch_domain.ElectricityBillFetchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.FetchBillPOST(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h *ElectricityBillFetchHandler) FetchBill(c echo.Context) error {
	consumerID := c.QueryParam("consumer_id")
	operatorCodeStr := c.QueryParam("operator_code")
	operatorCode, _ := strconv.Atoi(operatorCodeStr)

	res, err := h.Service.FetchBillGET(consumerID, operatorCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
