package service

import (
	"reflect"
	"testing"

	"bank-statement-viewer-backend/internal/model"
	"bank-statement-viewer-backend/internal/repository"

	"github.com/golang/mock/gomock"
)

func TestTransactionService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockTransactionRepository(ctrl)
	svc := NewTransactionService(mockRepo)

	t.Run("GetAll returns all transactions", func(t *testing.T) {
		expected := []model.Transaction{
			{Timestamp: 1, Name: "A", Type: "CREDIT", Amount: 100, Status: "SUCCESS"},
			{Timestamp: 2, Name: "B", Type: "DEBIT", Amount: 50, Status: "SUCCESS"},
		}
		mockRepo.EXPECT().GetAll().Return(expected)
		got := svc.GetAll()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("GetAll() = %v, want %v", got, expected)
		}
	})

	t.Run("SaveTransactions returns nil error for non-empty slice", func(t *testing.T) {
		txs := []model.Transaction{
			{Timestamp: 1, Name: "A", Type: "CREDIT", Amount: 100, Status: "SUCCESS"},
		}
		mockRepo.EXPECT().SaveAll(txs).Return(nil)
		err := svc.SaveTransactions(txs)
		if err != nil {
			t.Errorf("SaveTransactions() error = %v, want nil", err)
		}
	})

	t.Run("SaveTransactions returns error if nil slice", func(t *testing.T) {
		err := svc.SaveTransactions(nil)
		if err == nil {
			t.Errorf("SaveTransactions(nil) expected error, got nil")
		}
	})

	t.Run("CalculateBalance computes correct balance", func(t *testing.T) {
		transactions := []model.Transaction{
			{Timestamp: 1, Name: "A", Type: "CREDIT", Amount: 200, Status: "SUCCESS"},
			{Timestamp: 2, Name: "B", Type: "DEBIT", Amount: 50, Status: "SUCCESS"},
			{Timestamp: 3, Name: "C", Type: "DEBIT", Amount: 30, Status: "PENDING"}, // ignored
			{Timestamp: 4, Name: "D", Type: "CREDIT", Amount: 20, Status: "FAILED"}, // ignored
		}
		mockRepo.EXPECT().GetAll().Return(transactions)
		got := svc.CalculateBalance()
		want := int64(150) // 200 - 50
		if got != want {
			t.Errorf("CalculateBalance() = %v, want %v", got, want)
		}
	})

	t.Run("GetIssues returns only PENDING or FAILED transactions", func(t *testing.T) {
		transactions := []model.Transaction{
			{Timestamp: 1, Name: "A", Type: "CREDIT", Amount: 200, Status: "SUCCESS"},
			{Timestamp: 2, Name: "B", Type: "DEBIT", Amount: 50, Status: "FAILED"},
			{Timestamp: 3, Name: "C", Type: "DEBIT", Amount: 30, Status: "PENDING"},
		}
		mockRepo.EXPECT().GetAll().Return(transactions)
		got := svc.GetIssues()
		want := []model.Transaction{
			{Timestamp: 2, Name: "B", Type: "DEBIT", Amount: 50, Status: "FAILED"},
			{Timestamp: 3, Name: "C", Type: "DEBIT", Amount: 30, Status: "PENDING"},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetIssues() = %v, want %v", got, want)
		}
	})
}
