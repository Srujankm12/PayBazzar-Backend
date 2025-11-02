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


type WalletService struct {
	CreateRepo *moneytransfer_repository.WalletCreateRepo
	VerifyRepo *moneytransfer_repository.WalletVerifyRepo
	ApiBaseURL string
	ApiToken   string
}

// ---------- CONSTRUCTOR ----------

func NewWalletService(createRepo *moneytransfer_repository.WalletCreateRepo, verifyRepo *moneytransfer_repository.WalletVerifyRepo, baseURL, token string) *WalletService {
	return &WalletService{
		CreateRepo: createRepo,
		VerifyRepo: verifyRepo,
		ApiBaseURL: baseURL,
		ApiToken:   token,
	}
}
func (s *WalletService) CreateWalletPOST(req *moneytransfer_domain.CreateWalletRequest) (*moneytransfer_domain.CreateWalletResponse, error) {
	body, _ := json.Marshal(req)
	url := fmt.Sprintf("%s/moneytransfer/createWalletRequest", s.ApiBaseURL)

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %v", err)
	}
	defer resp.Body.Close()

	// 🔍 Log full HTTP details for debugging
	fmt.Println("---------------------------------------------------")
	fmt.Println("➡️ [CreateWallet] External API Request Sent")
	fmt.Println("URL:", url)
	fmt.Println("Headers:", httpReq.Header)
	fmt.Println("Request Body:", string(body))
	fmt.Println("HTTP Status:", resp.Status)

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(respBody))
	fmt.Println("---------------------------------------------------")

	var apiRes moneytransfer_domain.CreateWalletResponse
	if err := json.Unmarshal(respBody, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to parse API response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API error (HTTP %d): %s", resp.StatusCode, string(respBody))
	}

	return &apiRes, nil
}

func (s *WalletService) CreateWalletGET(mobileNo string) (*moneytransfer_domain.CreateWalletResponse, error) {
	url := fmt.Sprintf("%s/moneytransfer/createWalletRequest?mobile_no=%s", s.ApiBaseURL, mobileNo)
	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.CreateWalletResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	req := &moneytransfer_domain.CreateWalletRequest{MobileNo: mobileNo}
	if err := s.CreateRepo.SaveWalletRequest(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

// ---------- VERIFY WALLET OTP ----------

func (s *WalletService) VerifyOtpPOST(req *moneytransfer_domain.VerifyOtpRequest) (*moneytransfer_domain.VerifyOtpResponse, error) {
	url := s.ApiBaseURL + "/moneytransfer/verifyOtp"

	jsonData, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.VerifyOtpResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.VerifyRepo.SaveWalletVerification(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

func (s *WalletService) VerifyOtpGET(req *moneytransfer_domain.VerifyOtpRequest) (*moneytransfer_domain.VerifyOtpResponse, error) {
	url := fmt.Sprintf(
		"%s/moneytransfer/verifyOtp?mobile_no=%s&request_no=%s&otp=%s&firstName=%s&lastName=%s&addressLine1=%s&addressLine2=%s&city=%s&state=%s&pinCode=%s",
		s.ApiBaseURL, req.MobileNo, req.RequestNo, req.Otp, req.FirstName, req.LastName, req.AddressLine1,
		req.AddressLine2, req.City, req.State, req.PinCode,
	)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes moneytransfer_domain.VerifyOtpResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.VerifyRepo.SaveWalletVerification(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
