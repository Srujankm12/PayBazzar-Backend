package postpaidbill_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	postpaidbill_domain "github.com/paybazar-backend/internals/domain/BillPayments/postpaidBillFetch"
	postpaidbill_repository "github.com/paybazar-backend/internals/repository/BillPayments/postpaidBillFetch"
)

type PostpaidBillFetchService struct {
	Repo       *postpaidbill_repository.PostpaidBillFetchRepo
	ApiBaseURL string
	ApiToken   string
}

func NewPostpaidBillFetchService(repo *postpaidbill_repository.PostpaidBillFetchRepo, baseURL, token string) *PostpaidBillFetchService {
	return &PostpaidBillFetchService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *PostpaidBillFetchService) SavePostPaidBill(req *postpaidbill_domain.PostpaidBillRequest) (*postpaidbill_domain.PostpaidBillResponse, error) {
	body, _ := json.Marshal(req)
	url := s.ApiBaseURL + "/recharge/postPaidBillFetch"

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	var apiRes postpaidbill_domain.PostpaidBillResponse
	if err := json.Unmarshal(resBody, &apiRes); err != nil {
		return nil, err
	}
	if err := s.Repo.SavePostpaidBillFetch(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
func (s *PostpaidBillFetchService) FetchBill(req *postpaidbill_domain.PostpaidBillRequest) (*postpaidbill_domain.PostpaidBillResponse, error) {
	url := fmt.Sprintf("%s/recharge/postPaidBillFetch?mobile_no=%s&operator_code=%d",
		s.ApiBaseURL, req.MobileNo, req.OperatorCode)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var apiRes postpaidbill_domain.PostpaidBillResponse
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.SavePostpaidBillFetch(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}