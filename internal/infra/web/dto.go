// internal/infra/web/dto.go
package web

// EventRequestDTO representa o payload JSON esperado nas requisições POST /event
type EventRequestDTO struct {
	Type        string `json:"type"`
	Destination string `json:"destination,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      int    `json:"amount"`
}

// DepositResponseDTO representa a resposta JSON para operações de depósito
type DepositResponseDTO struct {
	Destination struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	} `json:"destination"`
}

// WithdrawResponseDTO representa a resposta JSON para operações de saque
type WithdrawResponseDTO struct {
	Origin struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	} `json:"origin"`
}

// TransferResponseDTO representa a resposta JSON para operações de transferência
type TransferResponseDTO struct {
	Origin struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	} `json:"origin"`
	Destination struct {
		ID      string `json:"id"`
		Balance int    `json:"balance"`
	} `json:"destination"`
}
