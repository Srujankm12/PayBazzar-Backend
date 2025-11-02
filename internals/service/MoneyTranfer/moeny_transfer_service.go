package moneytransfer_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	moneytransfer_domain "github.com/paybazar-backend/internals/domain/MoneyTransfer"
	moneytransfer_repository "github.com/paybazar-backend/internals/repository/MoneyTransfer"
)

type MoneyTransferService struct {
	Repo       *moneytransfer_repository.MoneyTransferRepository
	ApiBaseURL string
	ApiToken   string
}

func NewMoneyTransferService(repo *moneytransfer_repository.MoneyTransferRepository, baseURL, token string) *MoneyTransferService {
	return &MoneyTransferService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

// Initiate transfer (POST)
func (s *MoneyTransferService) InitiateTransfer(req *moneytransfer_domain.MRTransferRequest) (*moneytransfer_domain.MRTransferResponse, error) {
	url := s.ApiBaseURL + "/moneytransfer/mrTransfer"

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.MRTransferResponse
	if err := json.Unmarshal(respBody, &apiRes); err != nil {
		return nil, err
	}

	// save transaction
	if err := s.Repo.InsertMoneyTransfer(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

// Initiate transfer (GET version using query params) - optional helper
func (s *MoneyTransferService) InitiateTransferGET(req *moneytransfer_domain.MRTransferRequest) (*moneytransfer_domain.MRTransferResponse, error) {
	// build query
	url := fmt.Sprintf("%s/moneytransfer/mrTransfer?mobile_no=%s&beneficiary_name=%s&beneficiary_code=%s&partner_request_id=%s&amount=%v&account_no=%s&bankname=%s&ifsc=%s&transfer_type=%s",
		s.ApiBaseURL,
		req.MobileNo,
		req.BeneficiaryName,
		req.BeneficiaryCode,
		req.PartnerRequestID,
		req.Amount,
		req.AccountNo,
		req.BankName,
		req.IFSC,
		req.TransferType,
	)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.MRTransferResponse
	if err := json.Unmarshal(respBody, &apiRes); err != nil {
		return nil, err
	}

	// save transaction
	if err := s.Repo.InsertMoneyTransfer(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

// Check transfer status (POST)
func (s *MoneyTransferService) CheckStatus(req *moneytransfer_domain.TransferStatusRequest) (*moneytransfer_domain.TransferStatusResponse, error) {
	url := s.ApiBaseURL + "/moneytransfer/checkStatus"

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.TransferStatusResponse
	if err := json.Unmarshal(respBody, &apiRes); err != nil {
		return nil, err
	}

	// save status log
	if err := s.Repo.InsertTransferStatusLog(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

// Check transfer status (GET version)
func (s *MoneyTransferService) CheckStatusGET(transactionID string) (*moneytransfer_domain.TransferStatusResponse, error) {
	url := fmt.Sprintf("%s/moneytransfer/checkStatus?transaction_id=%s", s.ApiBaseURL, transactionID)
	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.TransferStatusResponse
	if err := json.Unmarshal(respBody, &apiRes); err != nil {
		return nil, err
	}

	// save status log
	req := &moneytransfer_domain.TransferStatusRequest{TransactionID: transactionID}
	if err := s.Repo.InsertTransferStatusLog(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
