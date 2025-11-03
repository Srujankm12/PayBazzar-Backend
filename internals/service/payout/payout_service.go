package payoutservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	payoutdomain "github.com/paybazar-backend/internals/domain/payout"
	payoutrepository "github.com/paybazar-backend/internals/repository/payout"
)

type PayoutService struct {
	Repo       *payoutrepository.PayoutRepo
	ApiBaseURL string
	ApiToken   string
	HTTPClient *http.Client
}

func NewPayoutService(repo *payoutrepository.PayoutRepo, baseURL, token string) *PayoutService {
	return &PayoutService{
		Repo:       repo,
		ApiBaseURL: baseURL,
		ApiToken:   token,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// -------------------- INITIATE PAYOUT --------------------
func (s *PayoutService) InitiatePayout(req *payoutdomain.PayoutRequest) (*payoutdomain.PayoutResponse, error) {
	url := fmt.Sprintf("%s/moneytransfer/mrTransfer", s.ApiBaseURL)

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	fmt.Printf("➡️ [Payout] Sending POST to: %s\n", url)

	resp, err := s.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("❌ failed to call RKIT payout API: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	raw := string(bodyBytes)
	fmt.Printf("📩 RKIT Response: %s\n", raw)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API error (HTTP %d): %s", resp.StatusCode, raw)
	}

	var apiRes payoutdomain.PayoutResponse
	if err := json.Unmarshal(bodyBytes, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to parse RKIT response: %v | raw: %s", err, raw)
	}

	// Store in DB
	if err := s.Repo.SavePayout(req, &apiRes); err != nil {
		return nil, fmt.Errorf("DB insert failed: %v", err)
	}

	return &apiRes, nil
}

// -------------------- CHECK STATUS --------------------
func (s *PayoutService) CheckPayoutStatus(req *payoutdomain.PayoutRequest) (*payoutdomain.PayoutResponse, error) {
	url := fmt.Sprintf("%s/moneytransfer/checkStatus", s.ApiBaseURL)

	body, _ := json.Marshal(req)
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Authorization", "Bearer "+s.ApiToken)
	httpReq.Header.Set("Content-Type", "application/json")

	fmt.Printf("➡️ [CheckStatus] Sending POST to: %s\n", url)

	resp, err := s.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("❌ failed to call RKIT status API: %v", err)
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	raw := string(bodyBytes)
	fmt.Printf("📩 RKIT Status Response: %s\n", raw)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external API error (HTTP %d): %s", resp.StatusCode, raw)
	}

	var apiRes payoutdomain.PayoutResponse
	if err := json.Unmarshal(bodyBytes, &apiRes); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %v | raw: %s", err, raw)
	}

	if err := s.Repo.SavePayout(req, &apiRes); err != nil {
		return nil, fmt.Errorf("DB insert failed: %v", err)
	}

	return &apiRes, nil
}
