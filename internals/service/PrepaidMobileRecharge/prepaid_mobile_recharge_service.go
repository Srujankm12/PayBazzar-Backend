package prepaidmobilerecharge_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	prepaidmobilerecharge_domain "github.com/paybazar-backend/internals/domain/PrepaidMobileRecharge"
	prepaidmobilerecharge_repository "github.com/paybazar-backend/internals/repository/PrepaidMobileRecharge"
)

type RechargeService struct {
	Repo       *prepaidmobilerecharge_repository.RechargeRepo
	ApiBaseURL string
	ApiToken   string
}

func NewRechargeService(repo *prepaidmobilerecharge_repository.RechargeRepo, baseURL, token string) *RechargeService {
	return &RechargeService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *RechargeService) DoRecharge(req *prepaidmobilerecharge_domain.RechargeRequest) (*prepaidmobilerecharge_domain.RechargeResponse, error) {
	jsonBody, _ := json.Marshal(req)
	url := fmt.Sprintf("%s/recharge/prepaid", s.ApiBaseURL)

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiRes prepaidmobilerecharge_domain.RechargeResponse
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return nil, err
	}
	if err := s.Repo.SaveRecharge(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
func (s *RechargeService) GetAllRecharges() ([]prepaidmobilerecharge_domain.RechargeResponse, error) {
	return s.Repo.GetAllRecharges()
}

func (s *RechargeService) GetRechargeByPartnerReqID(id string) (*prepaidmobilerecharge_domain.RechargeResponse, error) {
	return s.Repo.GetRechargeByPartnerReqID(id)
}
