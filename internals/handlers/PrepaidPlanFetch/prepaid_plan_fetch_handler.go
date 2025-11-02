package prepaidplanfetch_handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	prepaidplanfetch_domain "github.com/paybazar-backend/internals/domain/PrepaidPlanFetch"
	prepaidplanfetch_service "github.com/paybazar-backend/internals/service/PrepaidPlanFetch"
)

type PrepaidPlanHandler struct {
	Service *prepaidplanfetch_service.PrepaidPlanService
}

func NewPrepaidPlanHandler(s *prepaidplanfetch_service.PrepaidPlanService) *PrepaidPlanHandler {
	return &PrepaidPlanHandler{Service: s}
}

func (h *PrepaidPlanHandler) CreatePrepaidPlan(c echo.Context) error {
	var req prepaidplanfetch_domain.PrepaidPlanRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	res, err := h.Service.CreatePrepaidPlan(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *PrepaidPlanHandler) GetPrepaidPlan(c echo.Context) error {
	circle := c.QueryParam("circle")
	operator := c.QueryParam("operator_code")

	var req prepaidplanfetch_domain.PrepaidPlanRequest
	if circle == "" || operator == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Missing required query params"})
	}

	if _, err := fmt.Sscanf(circle, "%d", &req.Circle); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid circle"})
	}
	if _, err := fmt.Sscanf(operator, "%d", &req.OperatorCode); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid operator_code"})
	}

	res, err := h.Service.FetchPrepaidPlan(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
