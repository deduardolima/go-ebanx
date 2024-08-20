package web

import (
	"encoding/json"
	"go-ebanx/internal/usecase"
	"net/http"
	"strconv"
)

type Handler struct {
	GetBalanceUseCase *usecase.GetBalanceUseCase
	DepositUseCase    *usecase.DepositUseCase
	WithdrawUseCase   *usecase.WithdrawUseCase
	TransferUseCase   *usecase.TransferUseCase
}

func NewHandler(gbuc *usecase.GetBalanceUseCase, du *usecase.DepositUseCase, wu *usecase.WithdrawUseCase, tu *usecase.TransferUseCase) *Handler {
	return &Handler{
		GetBalanceUseCase: gbuc,
		DepositUseCase:    du,
		WithdrawUseCase:   wu,
		TransferUseCase:   tu,
	}
}

// @Summary Reseta o estado do sistema
// @Description Reseta o estado do sistema, deletando todas as contas.
// @ID reset-state
// @Tags Reset
// @Success 200 {object} ResetResponseDTO "OK"
// @Router /api/reset [post]
func (h *Handler) Reset(w http.ResponseWriter, r *http.Request) {
	h.DepositUseCase.AccountRepo.Reset()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// @Summary Obtém o saldo de uma conta bancária
// @Description Retorna o saldo da conta especificada
// @ID get-balance
// @Tags Operações
// @Param account_id query string true "ID da conta"
// @Success 200 {object} BalanceResponseDTO "Saldo da conta"
// @Failure 404 {object} ErrAccountNotFoundDTO "Conta não encontrada"
// @Failure 500 {object} ErrServerErrorDTO "Erro no servidor"
// @Router /api/balance [get]
func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")
	balance, err := h.GetBalanceUseCase.Execute(accountID)
	if err != nil {
		// For non-existing account, return 404 with just 0 in the body
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("0"))
		return
	}

	// For existing account, return 200 with just the balance in the body
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(balance)))
}

// @Summary Manipula eventos como depósito, saque e transferência
// @Description Manipula os eventos enviados para a API, realizando transações bancárias.
// @ID handle-event
// @Tags Operações
// @Accept json
// @Produce json
// @Param event body EventRequestDTO true "Evento bancário"
// @Success 201 {object} DepositResponseDTO "Depósito bem-sucedido"
// @Failure 404 {object} ErrAccountNotFoundDTO "Conta não encontrada"
// @Failure 500 {object} ErrServerErrorDTO "Erro no servidor"
// @Router /api/event [post]
func (h *Handler) Event(w http.ResponseWriter, r *http.Request) {
	var ev EventRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch ev.Type {
	case "deposit":
		h.handleDeposit(w, ev)
	case "withdraw":
		h.handleWithdraw(w, ev)
	case "transfer":
		h.handleTransfer(w, ev)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

// @Summary Depósito em conta
// @Description Realiza um depósito em uma conta existente.
// @ID handle-deposit
// @Tags Operações
// @Accept json
// @Produce json
// @Param event body EventRequestDTO true "Evento bancário"
// @Success 201 {object} DepositResponseDTO "Depósito bem-sucedido"
// @Failure 404 {object} ErrAccountNotFoundDTO "Conta não encontrada"
// @Failure 500 {object} ErrServerErrorDTO "Erro no servidor"
// @Router /api/event [post]
func (h *Handler) handleDeposit(w http.ResponseWriter, ev EventRequestDTO) {
	acc, err := h.DepositUseCase.Execute(ev.Destination, ev.Amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := DepositResponseDTO{}
	response.Destination.ID = acc.ID
	response.Destination.Balance = acc.Balance
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// @Summary Saque de conta
// @Description Realiza um saque de uma conta existente.
// @ID handle-withdraw
// @Tags Operações
// @Accept json
// @Produce json
// @Param event body EventRequestDTO true "Evento bancário"
// @Success 201 {object} WithdrawResponseDTO "Saque bem-sucedido"
// @Failure 404 {object} ErrAccountNotFoundDTO "Conta não encontrada"
// @Failure 500 {object} ErrServerErrorDTO "Erro no servidor"
// @Router /api/event [post]
func (h *Handler) handleWithdraw(w http.ResponseWriter, ev EventRequestDTO) {
	acc, err := h.WithdrawUseCase.Execute(ev.Origin, ev.Amount)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("0"))
		return
	}
	response := WithdrawResponseDTO{}
	response.Origin.ID = acc.ID
	response.Origin.Balance = acc.Balance
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// @Summary Transferência entre contas
// @Description Realiza uma transferência entre contas existentes.
// @ID handle-transfer
// @Tags Operações
// @Accept json
// @Produce json
// @Param event body EventRequestDTO true "Event"
// @Success 201 {object} TransferResponseDTO "Transferência bem-sucedida"
// @Failure 404 {object} ErrAccountNotFoundDTO "Conta não encontrada"
// @Failure 500 {object} ErrServerErrorDTO "Erro no servidor"
// @Router /api/event [post]
func (h *Handler) handleTransfer(w http.ResponseWriter, ev EventRequestDTO) {
	originAcc, destAcc, err := h.TransferUseCase.Execute(ev.Origin, ev.Destination, ev.Amount)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("0"))
		return
	}
	response := TransferResponseDTO{}
	response.Origin.ID = originAcc.ID
	response.Origin.Balance = originAcc.Balance
	response.Destination.ID = destAcc.ID
	response.Destination.Balance = destAcc.Balance
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
