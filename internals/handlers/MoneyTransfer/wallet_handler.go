package moneytransfer_handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
	moneytransfer_service "github.com/paybazar-backend/internals/service/MoneyTranfer"
)

type WalletHandler struct {
	Service *moneytransfer_service.WalletService
}

func NewWalletHandler(s *moneytransfer_service.WalletService) *WalletHandler {
	return &WalletHandler{Service: s}
}

// ---------------------- CREATE WALLET ----------------------
func (h *WalletHandler) CreateWalletPOST(c echo.Context) error {
	var req moneytransfer_domain.CreateWalletRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.CreateWalletPOST(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

// ---------------------- CHECK WALLET ----------------------
func (h *WalletHandler) CheckWalletExists(c echo.Context) error {
	mobileNo := c.QueryParam("mobile_no")
	if mobileNo == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "mobile_no is required"})
	}

	res, err := h.Service.CreateWalletGET(mobileNo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

// ---------------------- VERIFY OTP ----------------------
func (h *WalletHandler) VerifyOtpPOST(c echo.Context) error {
	var req moneytransfer_domain.VerifyOtpRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.VerifyOtpPOST(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

func (h *WalletHandler) VerifyOtpGET(c echo.Context) error {
	req := &moneytransfer_domain.VerifyOtpRequest{
		MobileNo:     c.QueryParam("mobile_no"),
		RequestNo:    c.QueryParam("request_no"),
		Otp:          c.QueryParam("otp"),
		FirstName:    c.QueryParam("firstName"),
		LastName:     c.QueryParam("lastName"),
		AddressLine1: c.QueryParam("addressLine1"),
		AddressLine2: c.QueryParam("addressLine2"),
		City:         c.QueryParam("city"),
		State:        c.QueryParam("state"),
		PinCode:      c.QueryParam("pinCode"),
	}

	if req.MobileNo == "" || req.RequestNo == "" || req.Otp == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Required fields missing"})
	}

	res, err := h.Service.VerifyOtpGET(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
