package prepaidplanfetch_domain

type PrepaidPlanRequest struct {
	OperatorCode int `json:"operator_code"`
	Circle       int `json:"circle"`
}

type PrepaidPlanResponse struct {
	Error  int    `json:"error"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	PlanData struct {
		CircleID int    `json:"circle_id"`
		ID       string `json:"_id"`
		Plan     []map[string][]struct {
			Amount      float64 `json:"amount"`
			Validity    string  `json:"validity"`
			Description string  `json:"description"`
			Talktime    string  `json:"talktime"`
			SMS         string  `json:"sms"`
			Disclaimer  string  `json:"disclaimer"`
			IsValid     int     `json:"is_valid"`
		} `json:"plan"`
	} `json:"planData"`
}
