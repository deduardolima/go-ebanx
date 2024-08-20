package web

// EventRequestDTO
// @Description
// @Success 201 {object} DepositResponseDTO
type EventRequestDTO struct {
	Type        string `json:"type" example:"deposit"`
	Destination string `json:"destination,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      int    `json:"amount" example:"100"`
}

// DepositResponseDTO
// @Description
type DepositResponseDTO struct {
	Destination struct {
		ID      string `json:"id" example:"100"`
		Balance int    `json:"balance" example:"200"`
	} `json:"destination"`
}

// WithdrawResponseDTO
// @Description
type WithdrawResponseDTO struct {
	Origin struct {
		ID      string `json:"id" example:"100"`
		Balance int    `json:"balance" example:"150"`
	} `json:"origin"`
}

// TransferResponseDTO
// @Description
type TransferResponseDTO struct {
	Origin struct {
		ID      string `json:"id" example:"100"`
		Balance int    `json:"balance" example:"50"`
	} `json:"origin"`
	Destination struct {
		ID      string `json:"id" example:"200"`
		Balance int    `json:"balance" example:"150"`
	} `json:"destination"`
}

// BalanceResponseDTO
// @Description
type BalanceResponseDTO struct {
	Balance int `json:"balance" example:"200"`
}

// ResetResponseDTO
// @Description
type ResetResponseDTO struct {
	Status string `json:"status" example:"OK"`
}

// ErrAccountNotFoundDTO
// @Description
type ErrAccountNotFoundDTO struct {
	NotFound string `json:"not_found" example:"account not found"`
}

// ErrServerErrorDTO
// @Description
type ErrServerErrorDTO struct {
	Error string `json:"error" example:"internal server error"`
}
