package ott_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	ott_domain "github.com/paybazar-backend/internals/domain/Ott"
	ott_repository "github.com/paybazar-backend/internals/repository/Ott"
)

type OTTSubscriptionService struct {
	Repo       *ott_repository.OTTSubscriptionRepo
	ApiBaseURL string
	ApiToken   string
}

func NewOTTSubscriptionService(repo *ott_repository.OTTSubscriptionRepo, baseURL, token string) *OTTSubscriptionService {
	return &OTTSubscriptionService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *OTTSubscriptionService) CreateSubscription(req *ott_domain.OTTSubscriptionRequest) (*ott_domain.OTTSubscriptionResponse, error) {
	body, _ := json.Marshal(req)
	url := s.ApiBaseURL + "/recharge/ott"

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
	var apiRes ott_domain.OTTSubscriptionResponse
	if err := json.Unmarshal(resBody, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.SaveOTTSubscription(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}

func (s *OTTSubscriptionService) CreateSubscriptionGET(req *ott_domain.OTTSubscriptionRequest) (*ott_domain.OTTSubscriptionResponse, error) {
	url := fmt.Sprintf("%s/recharge/ott?mobile_no=%s&operator_code=%d&amount=%.2f&partner_request_id=%s&customer_email=%s&user_var1=%s&user_var2=%s&user_var3=%s",
		s.ApiBaseURL, req.MobileNo, req.OperatorCode, req.Amount, req.PartnerRequestID,
		req.CustomerEmail, req.UserVar1, req.UserVar2, req.UserVar3,
	)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	var apiRes ott_domain.OTTSubscriptionResponse
	if err := json.Unmarshal(resBody, &apiRes); err != nil {
		return nil, err
	}

	if err := s.Repo.SaveOTTSubscription(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
