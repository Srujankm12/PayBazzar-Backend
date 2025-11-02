package moneytransfer_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
	moneytransfer_service "github.com/paybazar-backend/internals/service/MoneyTranfer"
)

type BeneficiaryHandler struct {
	Service *moneytransfer_service.BeneficiaryService
}

func NewBeneficiaryHandler(service *moneytransfer_service.BeneficiaryService) *BeneficiaryHandler {
	return &BeneficiaryHandler{Service: service}
}

// ---------------------- ADD BENEFICIARY ----------------------
func (h *BeneficiaryHandler) AddBeneficiaryHandler(c echo.Context) error {
	var req moneytransfer_domain.AddBeneficiaryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	resp, err := h.Service.AddBeneficiary(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// ---------------------- GET BENEFICIARIES ----------------------
func (h *BeneficiaryHandler) GetBeneficiariesHandler(c echo.Context) error {
	mobileNo := c.QueryParam("mobile_no")
	if mobileNo == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "mobile_no is required"})
	}

	resp, err := h.Service.GetBeneficiaries(mobileNo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// ---------------------- DELETE BENEFICIARY ----------------------
func (h *BeneficiaryHandler) DeleteBeneficiaryHandler(c echo.Context) error {
	var req moneytransfer_domain.DeleteBeneficiaryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	resp, err := h.Service.DeleteBeneficiary(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// ---------------------- VERIFY DELETE BENEFICIARY ----------------------
func (h *BeneficiaryHandler) VerifyDeleteBeneficiaryHandler(c echo.Context) error {
	var req moneytransfer_domain.VerifyDeleteBeneficiaryRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	resp, err := h.Service.VerifyDeleteBeneficiary(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
