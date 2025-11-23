package handler

import (
	"io"
	"net/http"

	"bank-statement-viewer-backend/internal/service"
	"bank-statement-viewer-backend/internal/utils"
	"bank-statement-viewer-backend/pkg/response"
)

type TransactionHandler struct {
	svc service.TransactionService
}

func NewTransactionHandler(s service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: s}
}

// POST /upload
// form field "file" expected
func (h *TransactionHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.JSON(w, http.StatusMethodNotAllowed, response.NewError("method not allowed"))
		return
	}
	// limit size to e.g. 10MB
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

	// read the file into parser
	transactions, err := utils.ParseCSVTransactions(file)
	if err != nil {
		// try to read file as text and return partial contents (helpful for debugging)
		body, _ := io.ReadAll(file)
		_ = body
		response.JSON(w, http.StatusBadRequest, response.NewError("failed to parse csv: "+err.Error()))
		return
	}

	if err := h.svc.SaveTransactions(transactions); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.NewError("failed to save transactions: "+err.Error()))
		return
	}

	response.JSON(w, http.StatusOK, response.NewData(map[string]interface{}{
		"count": len(transactions),
	}))
}

// GET /balance
func (h *TransactionHandler) BalanceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, response.NewError("method not allowed"))
		return
	}
	b := h.svc.CalculateBalance()
	response.JSON(w, http.StatusOK, response.NewData(map[string]int64{"balance": b}))
}

// GET /issues
func (h *TransactionHandler) IssuesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.JSON(w, http.StatusMethodNotAllowed, response.NewError("method not allowed"))
		return
	}
	issues := h.svc.GetIssues()
	// encode result
	response.JSON(w, http.StatusOK, response.NewData(issues))
}

// small helper to read raw body for debugging (not used)
func readAll(r io.Reader) []byte {
	b, _ := io.ReadAll(r)
	return b
}
