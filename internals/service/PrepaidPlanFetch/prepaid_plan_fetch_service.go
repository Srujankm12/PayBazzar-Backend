package prepaidplanfetch_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	prepaidplanfetch_domain "github.com/paybazar-backend/internals/domain/PrepaidPlanFetch"
	prepaidplanfetch_repository "github.com/paybazar-backend/internals/repository/PrepaidPlanFetch"
)

type PrepaidPlanService struct {
	Repo       *prepaidplanfetch_repository.PrepaidPlanRepository
	ApiBaseURL string
	ApiToken   string
}

func NewPrepaidPlanService(repo *prepaidplanfetch_repository.PrepaidPlanRepository, baseURL, token string) *PrepaidPlanService {
	return &PrepaidPlanService{
		Repo:       repo,
		ApiBaseURL: baseURL,
		ApiToken:   token,
	}
}

// 🔸 POST version
func (s *PrepaidPlanService) CreatePrepaidPlan(req *prepaidplanfetch_domain.PrepaidPlanRequest) (*prepaidplanfetch_domain.PrepaidPlanResponse, error) {
	body, _ := json.Marshal(req)
	url := fmt.Sprintf("%s/recharge/prepaidPlanFetch", s.ApiBaseURL)

	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	raw := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API error (HTTP %d): %s", resp.StatusCode, raw)
	}

	var apiRes prepaidplanfetch_domain.PrepaidPlanResponse
	if err := json.Unmarshal(bodyBytes, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v | raw: %s", err, raw)
	}

	if err := s.Repo.SavePlan(req, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to save plan: %v", err)
	}

	return &apiRes, nil
}

// 🔸 GET version
func (s *PrepaidPlanService) FetchPrepaidPlan(req *prepaidplanfetch_domain.PrepaidPlanRequest) (*prepaidplanfetch_domain.PrepaidPlanResponse, error) {
	url := fmt.Sprintf("%s/recharge/prepaidPlanFetch?circle=%d&operator_code=%d",
		s.ApiBaseURL, req.Circle, req.OperatorCode)

	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	raw := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API error (HTTP %d): %s", resp.StatusCode, raw)
	}

	var apiRes prepaidplanfetch_domain.PrepaidPlanResponse
	if err := json.Unmarshal(bodyBytes, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v | raw: %s", err, raw)
	}

	if err := s.Repo.SavePlan(req, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to save plan: %v", err)
	}

	return &apiRes, nil
}
