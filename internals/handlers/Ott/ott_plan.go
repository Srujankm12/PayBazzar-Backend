package ott_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	ott_domain "github.com/paybazar-backend/internals/domain/Ott"
	ott_service "github.com/paybazar-backend/internals/service/Ott"
)

type OTTPlanHandler struct {
	Service *ott_service.OTTPlanService
}

func NewOTTPlanHandler(s *ott_service.OTTPlanService) *OTTPlanHandler {
	return &OTTPlanHandler{Service: s}
}

func (h *OTTPlanHandler) FetchOTTPlans(c echo.Context) error {
	operatorCodeStr := c.QueryParam("operator_code")
	if operatorCodeStr == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing operator_code"})
	}

	operatorCode, err := strconv.Atoi(operatorCodeStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid operator_code"})
	}

	req := ott_domain.OTTPlanRequest{OperatorCode: operatorCode}
	res, err := h.Service.FetchOTTPlans(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
