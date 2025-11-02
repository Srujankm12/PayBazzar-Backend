package moneytransfer_domain

// ------------------ ADD BENEFICIARY ------------------
type AddBeneficiaryRequest struct {
	MobileNo        string `json:"mobile_no"`
	BeneficiaryName string `json:"beneficiary_name"`
	BankName        string `json:"bank_name"`
	AccountNo       string `json:"account_no"`
	IFSC            string `json:"ifsc"`
}

type AddBeneficiaryResponse struct {
	Error         int    `json:"error"`
	Msg           string `json:"msg"`
	BeneficiaryID string `json:"beneficiary_id"`
	Description   string `json:"description"`
}

// ------------------ GET BENEFICIARY ------------------
type GetBeneficiaryResponse struct {
	Error          int    `json:"error"`
	Msg            string `json:"msg"`
	Description    string `json:"description"`
	BeneficiaryList []struct {
		BeneficiaryID string `json:"beneficiary_id"`
		UUID          string `json:"uuid"`
		AccountDetail struct {
			AccountNumber     string `json:"account_number"`
			IFSCCode          string `json:"ifsc_code"`
			BankName          string `json:"bank_name"`
			AccountHolderName string `json:"account_holder_name"`
		} `json:"account_detail"`
	} `json:"beneficiary_list"`
}

// ------------------ DELETE BENEFICIARY ------------------
type DeleteBeneficiaryRequest struct {
	MobileNo      string `json:"mobile_no"`
	BeneficiaryID string `json:"beneficiary_id"`
}

type DeleteBeneficiaryResponse struct {
	Error       int    `json:"error"`
	Msg         string `json:"msg"`
	RequestNo   string `json:"request_no"`
	Description string `json:"description"`
}

// ------------------ VERIFY DELETE BENEFICIARY ------------------
type VerifyDeleteBeneficiaryRequest struct {
	MobileNo      string `json:"mobile_no"`
	RequestNo     string `json:"request_no"`
	OTP           string `json:"otp"`
	BeneficiaryID string `json:"beneficiary_id"`
}

type VerifyDeleteBeneficiaryResponse struct {
	Error       int    `json:"error"`
	Msg         string `json:"msg"`
	Description string `json:"description"`
}
