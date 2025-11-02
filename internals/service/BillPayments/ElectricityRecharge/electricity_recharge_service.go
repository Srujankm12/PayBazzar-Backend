package electricitybillpayment_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	electricitybillpayment_domain "github.com/paybazar-backend/internals/domain/BillPayments/ElectricityRecharge"
	electricitybillpayment_repository "github.com/paybazar-backend/internals/repository/BillPayments/ElectricityRecharge"
)

type ElectricityBillPaymentService struct {
	Repo       *electricitybillpayment_repository.ElectricityBillPaymentRepo
	ApiBaseURL string
	ApiToken   string
}

func NewElectricityBillPaymentService(repo *electricitybillpayment_repository.ElectricityBillPaymentRepo, baseURL, token string) *ElectricityBillPaymentService {
	return &ElectricityBillPaymentService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *ElectricityBillPaymentService) MakePaymentPOST(req *electricitybillpayment_domain.ElectricityBillPaymentRequest) (*electricitybillpayment_domain.ElectricityBillPaymentResponse, error) {
	body, _ := json.Marshal(req)
	url := fmt.Sprintf("%s/recharge/billpayment?", s.ApiBaseURL)

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes electricitybillpayment_domain.ElectricityBillPaymentResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	_ = s.Repo.SavePayment(req, &apiRes)
	return &apiRes, nil
}

func (s *ElectricityBillPaymentService) MakePaymentGET(req *electricitybillpayment_domain.ElectricityBillPaymentRequest) (*electricitybillpayment_domain.ElectricityBillPaymentResponse, error) {
	url := fmt.Sprintf("%s/recharge/billpayment?p1=%s&p2=%s&p3=%s&customer_email=%s&operator_code=%d&amount=%f&partner_request_id=%s&user_var1=%s&user_var2=%s&user_var3=%s",
		s.ApiBaseURL, req.P1, req.P2, req.P3, req.CustomerEmail, req.OperatorCode, req.Amount, req.PartnerRequestID,
		req.UserVar1, req.UserVar2, req.UserVar3,
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
	var apiRes electricitybillpayment_domain.ElectricityBillPaymentResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	_ = s.Repo.SavePayment(req, &apiRes)
	return &apiRes, nil
}
