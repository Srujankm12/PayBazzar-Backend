package prepaidmobilerecharge_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	prepaidmobilerecharge_domain "github.com/paybazar-backend/internals/domain/PrepaidMobileRecharge"
	prepaidmobilerecharge_service "github.com/paybazar-backend/internals/service/PrepaidMobileRecharge"
)

type RechargeHandler struct {
	Service *prepaidmobilerecharge_service.RechargeService
}

func NewRechargeHandler(service *prepaidmobilerecharge_service.RechargeService) *RechargeHandler {
	return &RechargeHandler{Service: service}
}

func (h *RechargeHandler) Recharge(c echo.Context) error {
	var req prepaidmobilerecharge_domain.RechargeRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	res, err := h.Service.DoRecharge(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
func (h *RechargeHandler) GetAllRecharges(c echo.Context) error {
	recharges, err := h.Service.GetAllRecharges()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, recharges)
}

func (h *RechargeHandler) GetRechargeByPartnerReqID(c echo.Context) error {
	id := c.Param("partner_request_id")
	recharge, err := h.Service.GetRechargeByPartnerReqID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Recharge not found"})
	}
	return c.JSON(http.StatusOK, recharge)
}
