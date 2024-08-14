// internal/infra/web/handler.go
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

func (h *Handler) Reset(w http.ResponseWriter, r *http.Request) {
	h.DepositUseCase.AccountRepo.Reset()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountID := r.URL.Query().Get("account_id")
	balance, err := h.GetBalanceUseCase.Execute(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("0"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(balance)))
}

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
