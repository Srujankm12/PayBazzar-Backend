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

type BeneficiaryService struct {
	Repo       *moneytransfer_repository.BeneficiaryRepository
	ApiBaseURL string
	ApiToken   string
}

func NewBeneficiaryService(repo *moneytransfer_repository.BeneficiaryRepository, baseURL, token string) *BeneficiaryService {
	return &BeneficiaryService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

// ---------------------- ADD BENEFICIARY ----------------------
func (s *BeneficiaryService) AddBeneficiary(req *moneytransfer_domain.AddBeneficiaryRequest) (*moneytransfer_domain.AddBeneficiaryResponse, error) {
	url := s.ApiBaseURL + "/moneytransfer/addBeneficiaryRequest"

	data, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.AddBeneficiaryResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.InsertAddBeneficiary(req, &apiRes); err != nil {
		return nil, err
	}
	return &apiRes, nil
}

// ---------------------- GET BENEFICIARIES ----------------------
func (s *BeneficiaryService) GetBeneficiaries(mobileNo string) (*moneytransfer_domain.GetBeneficiaryResponse, error) {
	url := fmt.Sprintf("%s/moneytransfer/getUserDetails?mobile_no=%s", s.ApiBaseURL, mobileNo)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.GetBeneficiaryResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.InsertBeneficiaryList(mobileNo, &apiRes); err != nil {
		return nil, err
	}
	return &apiRes, nil
}

// ---------------------- DELETE BENEFICIARY ----------------------
func (s *BeneficiaryService) DeleteBeneficiary(req *moneytransfer_domain.DeleteBeneficiaryRequest) (*moneytransfer_domain.DeleteBeneficiaryResponse, error) {
	url := s.ApiBaseURL + "/moneytransfer/deleteBeneficiaryRequest"

	data, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.DeleteBeneficiaryResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.InsertDeleteBeneficiary(req, &apiRes); err != nil {
		return nil, err
	}
	return &apiRes, nil
}

// ---------------------- VERIFY DELETE BENEFICIARY ----------------------
func (s *BeneficiaryService) VerifyDeleteBeneficiary(req *moneytransfer_domain.VerifyDeleteBeneficiaryRequest) (*moneytransfer_domain.VerifyDeleteBeneficiaryResponse, error) {
	url := s.ApiBaseURL + "/moneytransfer/confirmDeleteBeneficiary"

	data, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.VerifyDeleteBeneficiaryResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.InsertVerifyDeleteBeneficiary(req, &apiRes); err != nil {
		return nil, err
	}
	return &apiRes, nil
}
