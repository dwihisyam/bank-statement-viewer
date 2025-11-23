package handler

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"bank-statement-viewer-backend/internal/model"
	"bank-statement-viewer-backend/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := service.NewMockTransactionService(ctrl)
	handler := NewTransactionHandler(mockService)

	t.Run("BalanceHandler returns balance", func(t *testing.T) {
		mockService.EXPECT().CalculateBalance().Return(int64(1000))
		req := httptest.NewRequest(http.MethodGet, "/balance", nil)
		w := httptest.NewRecorder()

		handler.BalanceHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, string(body), `"balance":1000`)
	})

	t.Run("BalanceHandler wrong method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/balance", nil)
		w := httptest.NewRecorder()

		handler.BalanceHandler(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	})

	t.Run("IssuesHandler returns issues", func(t *testing.T) {
		issues := []model.Transaction{
			{Timestamp: 1, Name: "A", Type: "DEBIT", Amount: 50, Status: "FAILED"},
		}
		mockService.EXPECT().GetIssues().Return(issues)
		req := httptest.NewRequest(http.MethodGet, "/issues", nil)
		w := httptest.NewRecorder()

		handler.IssuesHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, string(body), `"name":"A"`)
	})

	t.Run("IssuesHandler wrong method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/issues", nil)
		w := httptest.NewRecorder()

		handler.IssuesHandler(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	})

	t.Run("UploadHandler successfully saves transactions", func(t *testing.T) {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)
		fw, _ := writer.CreateFormFile("file", "transactions.csv")
		_, _ = fw.Write([]byte("1,A,CREDIT,100,SUCCESS,test\n"))
		writer.Close()

		mockService.EXPECT().SaveTransactions(gomock.Any()).Return(nil)
		req := httptest.NewRequest(http.MethodPost, "/upload", &b)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handler.UploadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Contains(t, string(body), `"count":1`)
	})

	t.Run("UploadHandler wrong method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/upload", nil)
		w := httptest.NewRecorder()

		handler.UploadHandler(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
	})

	t.Run("UploadHandler missing file", func(t *testing.T) {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)
		writer.Close()
		req := httptest.NewRequest(http.MethodPost, "/upload", &b)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handler.UploadHandler(w, req)

		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, readBody(res), "file is required")
	})

	t.Run("UploadHandler service returns error", func(t *testing.T) {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)
		fw, _ := writer.CreateFormFile("file", "transactions.csv")
		_, _ = fw.Write([]byte("1,A,CREDIT,100,SUCCESS,test\n"))
		writer.Close()

		mockService.EXPECT().SaveTransactions(gomock.Any()).Return(errors.New("failed to save"))
		req := httptest.NewRequest(http.MethodPost, "/upload", &b)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handler.UploadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
		assert.Contains(t, string(body), "failed to save")
	})

	t.Run("UploadHandler CSV parse error", func(t *testing.T) {
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)
		fw, _ := writer.CreateFormFile("file", "transactions.csv")
		// CSV invalid (misal timestamp bukan int)
		_, _ = fw.Write([]byte("INVALID,CREDIT,100,SUCCESS,test\n"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/upload", &b)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handler.UploadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, string(body), "failed to parse csv")
	})

	t.Run("UploadHandler parse multipart form error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockService := service.NewMockTransactionService(ctrl)
		handler := NewTransactionHandler(mockService)

		// Buat body yang sangat besar supaya melebihi 10MB limit
		largeBody := make([]byte, 11<<20) // 11 MB
		req := httptest.NewRequest(http.MethodPost, "/upload", bytes.NewReader(largeBody))
		req.Header.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")

		w := httptest.NewRecorder()
		handler.UploadHandler(w, req)

		res := w.Result()
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Contains(t, string(body), "failed to parse form")
	})


}

func readBody(res *http.Response) string {
	body, _ := io.ReadAll(res.Body)
	return string(body)
}
