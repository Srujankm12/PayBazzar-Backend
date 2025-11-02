package postpaidmobilerecharge_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	postpaidmobilerecharge_domain "github.com/paybazar-backend/internals/domain/BillPayments/postpaidmobilerecharge"
	postpaidmobilerecharge_repository "github.com/paybazar-backend/internals/repository/BillPayments/postpaidmobilerecharge"
)

type PostpaidRechargeService struct {
	Repo       *postpaidmobilerecharge_repository.PostpaidMobileRechargeRepo
	ApiBaseURL string
	ApiToken   string
}

func NewPostpaidRechargeService(repo *postpaidmobilerecharge_repository.PostpaidMobileRechargeRepo, baseURL, token string) *PostpaidRechargeService {
	return &PostpaidRechargeService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

// ✅ POST Method
func (s *PostpaidRechargeService) ProcessPostpaidRecharge(req *postpaidmobilerecharge_domain.PostpaidMobileRechargeRequest) (*postpaidmobilerecharge_domain.PostpaidMobileRechargeResponse, error) {
	body, _ := json.Marshal(req)
	url := s.ApiBaseURL + "/recharge/postpaid"

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
	var apiRes postpaidmobilerecharge_domain.PostpaidMobileRechargeResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.SavePostpaidRecharge(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

// ✅ GET Method
func (s *PostpaidRechargeService) FetchPostpaidRecharge(req *postpaidmobilerecharge_domain.PostpaidMobileRechargeRequest) (*postpaidmobilerecharge_domain.PostpaidMobileRechargeResponse, error) {

	url := fmt.Sprintf(
		"%s/recharge/postpaid?mobile_no=%s&operator_code=%d&amount=%v&partner_request_id=%s&circle=%d&recharge_type=%d&user_var1=%s&user_var2=%s&user_var3=%s",
		s.ApiBaseURL,
		req.MobileNo,
		req.OperatorCode,
		req.Amount,
		req.PartnerRequestID,
		req.Circle,
		req.RechargeType,
		req.UserVar1,
		req.UserVar2,
		req.UserVar3,
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
	var apiRes postpaidmobilerecharge_domain.PostpaidMobileRechargeResponse
	if err := json.Unmarshal(bodyRes, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.SavePostpaidRecharge(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
