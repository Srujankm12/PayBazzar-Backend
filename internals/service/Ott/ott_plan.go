package ott_service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	ott_domain "github.com/paybazar-backend/internals/domain/Ott"
	ott_repository "github.com/paybazar-backend/internals/repository/Ott"
)

type OTTPlanService struct {
	Repo       *ott_repository.OTTPlanRepo
	ApiBaseURL string
	ApiToken   string
}

func NewOTTPlanService(repo *ott_repository.OTTPlanRepo, baseURL, token string) *OTTPlanService {
	return &OTTPlanService{Repo: repo, ApiBaseURL: baseURL, ApiToken: token}
}

func (s *OTTPlanService) FetchOTTPlans(req *ott_domain.OTTPlanRequest) (*ott_domain.OTTPlanResponse, error) {
	url := fmt.Sprintf("%s/recharge/getottplan?operator_code=%d", s.ApiBaseURL, req.OperatorCode)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var apiRes ott_domain.OTTPlanResponse
	if err := json.Unmarshal(body, &apiRes); err != nil {
		return nil, err
	}
	if err := s.Repo.SaveOTTPlan(req, &apiRes); err != nil {
		return nil, err
	}

	return &apiRes, nil
}
