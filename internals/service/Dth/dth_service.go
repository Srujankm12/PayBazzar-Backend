package dth_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dth_domain "github.com/paybazar-backend/internals/domain/Dth"
	dth_repository "github.com/paybazar-backend/internals/repository/Dth"
)

type DTHRechargeService struct {
	Repo       *dth_repository.DTHRechargeRepo
	ApiBaseURL string
	ApiToken   string
}

func NewDTHRechargeService(repo *dth_repository.DTHRechargeRepo, baseURL, token string) *DTHRechargeService {
	return &DTHRechargeService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *DTHRechargeService) ProcessDTHRecharge(req *dth_domain.DTHRechargeRequest) (*dth_domain.DTHRechargeResponse, error) {
	body, _ := json.Marshal(req)
	url := s.ApiBaseURL + "/recharge/dth"

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
	var apiRes dth_domain.DTHRechargeResponse
	if err := json.Unmarshal(resBody, &apiRes); err != nil {
		return nil, err
	}
	if err := s.Repo.SaveDTHRecharge(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
func (s *DTHRechargeService) FetchDTHRecharge(req *dth_domain.DTHRechargeRequest) (*dth_domain.DTHRechargeResponse, error) {
	url := s.ApiBaseURL + "/recharge/dth" +
		"?customer_id=" + req.CustomerID +
		"&operator_code=" + fmt.Sprintf("%d", req.OperatorCode) +
		"&amount=" + fmt.Sprintf("%v", req.Amount) +
		"&partner_request_id=" + req.PartnerRequestID +
		"&user_var1=" + req.UserVar1 +
		"&user_var2=" + req.UserVar2 +
		"&user_var3=" + req.UserVar3

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)
	var apiRes dth_domain.DTHRechargeResponse
	if err := json.Unmarshal(resBody, &apiRes); err != nil {
		return nil, err
	}
	if err := s.Repo.SaveDTHRecharge(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
