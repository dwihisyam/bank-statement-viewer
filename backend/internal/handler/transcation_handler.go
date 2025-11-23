package handler

import (
	"io"
	"net/http"

	"bank-statement-viewer-backend/internal/service"
	"bank-statement-viewer-backend/internal/utils"
	"bank-statement-viewer-backend/pkg/response"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, response.NewError("method not allowed"))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		response.JSON(w, http.StatusBadRequest, response.NewError("failed to parse form: "+err.Error()))
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.NewError("file is required: "+err.Error()))
		return
	}
	defer file.Close()

	transactions, err := utils.ParseCSVTransactions(file)
	if err != nil {
		body, _ := io.ReadAll(file)
		_ = body
		response.JSON(w, http.StatusBadRequest, response.NewError("failed to parse csv: "+err.Error()))
		return
	}

	if err := h.service.SaveTransactions(transactions); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.NewError("failed to save transactions: "+err.Error()))
		return
	}

	response.JSON(w, http.StatusOK, response.NewData(map[string]interface{}{
		"count": len(transactions),
	}))
}

func (h *TransactionHandler) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, response.NewError("method not allowed"))
		return
	}
	b := h.service.CalculateBalance()
	response.JSON(w, http.StatusOK, response.NewData(map[string]int64{"balance": b}))
}

func (h *TransactionHandler) IssuesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, response.NewError("method not allowed"))
		return
	}
	issues := h.service.GetIssues()
	response.JSON(w, http.StatusOK, response.NewData(issues))
}
