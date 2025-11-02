package moneytransfer_handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
	moneytransfer_service "github.com/paybazar-backend/internals/service/MoneyTranfer"
)

type MoneyTransferHandler struct {
	Service *moneytransfer_service.MoneyTransferService
}

func NewMoneyTransferHandler(s *moneytransfer_service.MoneyTransferService) *MoneyTransferHandler {
	return &MoneyTransferHandler{Service: s}
}

// ---------------------- INITIATE TRANSFER ----------------------
func (h *MoneyTransferHandler) InitiateTransferPOST(c echo.Context) error {
	var req moneytransfer_domain.MRTransferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.InitiateTransfer(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

// ---------------------- INITIATE TRANSFER (GET) ----------------------
func (h *MoneyTransferHandler) InitiateTransferGET(c echo.Context) error {
	amountStr := c.QueryParam("amount")
	amount, _ := strconv.ParseFloat(amountStr, 64)

	req := &moneytransfer_domain.MRTransferRequest{
		MobileNo:         c.QueryParam("mobile_no"),
		BeneficiaryName:  c.QueryParam("beneficiary_name"),
		BeneficiaryCode:  c.QueryParam("beneficiary_code"),
		PartnerRequestID: c.QueryParam("partner_request_id"),
		Amount:           amount,
		AccountNo:        c.QueryParam("account_no"),
		BankName:         c.QueryParam("bankname"),
		IFSC:             c.QueryParam("ifsc"),
		TransferType:     c.QueryParam("transfer_type"),
		UserVar1:         c.QueryParam("user_var1"),
		UserVar2:         c.QueryParam("user_var2"),
		UserVar3:         c.QueryParam("user_var3"),
	}

	res, err := h.Service.InitiateTransferGET(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

// ---------------------- CHECK STATUS ----------------------
func (h *MoneyTransferHandler) CheckStatusPOST(c echo.Context) error {
	var req moneytransfer_domain.TransferStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	res, err := h.Service.CheckStatus(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}

// ---------------------- CHECK STATUS (GET) ----------------------
func (h *MoneyTransferHandler) CheckStatusGET(c echo.Context) error {
	transactionID := c.QueryParam("transaction_id")
	if transactionID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "transaction_id is required"})
	}

	res, err := h.Service.CheckStatusGET(transactionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
