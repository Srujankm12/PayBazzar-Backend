package electricitybillfetch_service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	electricitybillfetch_domain "github.com/paybazar-backend/internals/domain/BillPayments/ElectricityBill"
	electricitybillfetch_repository "github.com/paybazar-backend/internals/repository/BillPayments/ElectricityBill"
)

type ElectricityBillFetchService struct {
	Repo       *electricitybillfetch_repository.ElectricityBillFetchRepo
	ApiBaseURL string
	ApiToken   string
}

func NewElectricityBillFetchService(repo *electricitybillfetch_repository.ElectricityBillFetchRepo, baseURL, token string) *ElectricityBillFetchService {
	return &ElectricityBillFetchService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *ElectricityBillFetchService) FetchBillPOST(req *electricitybillfetch_domain.ElectricityBillFetchRequest) (*electricitybillfetch_domain.ElectricityBillFetchResponse, error) {
	body, _ := json.Marshal(req)
	url := s.ApiBaseURL + "/recharge/electricityBillFetch"

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
	var apiRes electricitybillfetch_domain.ElectricityBillFetchResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	_ = s.Repo.SaveElectricityBillFetch(req, &apiRes)
	return &apiRes, nil
}

func (s *ElectricityBillFetchService) FetchBillGET(consumerID string, operatorCode int) (*electricitybillfetch_domain.ElectricityBillFetchResponse, error) {
	url := s.ApiBaseURL + "/recharge/electricityBillFetch?consumer_id=" + consumerID + "&operator_code=" + string(rune(operatorCode))

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyRes, _ := io.ReadAll(resp.Body)
	var apiRes electricitybillfetch_domain.ElectricityBillFetchResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	req := &electricitybillfetch_domain.ElectricityBillFetchRequest{
		ConsumerID:   consumerID,
		OperatorCode: operatorCode,
	}
	_ = s.Repo.SaveElectricityBillFetch(req, &apiRes)

	return &apiRes, nil
}
