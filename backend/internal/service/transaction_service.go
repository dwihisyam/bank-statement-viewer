package service

import (
	"errors"

	"bank-statement-viewer-backend/internal/model"
	"bank-statement-viewer-backend/internal/repository"
)

type TransactionService interface {
	SaveTransactions([]model.Transaction) error
	CalculateBalance() int64
	GetIssues() []model.Transaction
	GetAll() []model.Transaction
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{repo: r}
}

func (s *transactionService) SaveTransactions(tx []model.Transaction) error {
	if tx == nil {
		return errors.New("no transactions")
	}
	// Save as replacement (behaviour can be changed to AppendAll)
	return s.repo.SaveAll(tx)
}

func (s *transactionService) GetAll() []model.Transaction {
	return s.repo.GetAll()
}

// CalculateBalance returns credits - debits considering only SUCCESS status
func (s *transactionService) CalculateBalance() int64 {
	all := s.repo.GetAll()
	var credits int64
	var debits int64
	for _, t := range all {
		if t.Status != "SUCCESS" {
			continue
		}
		if t.Type == "CREDIT" {
			credits += t.Amount
		} else if t.Type == "DEBIT" {
			debits += t.Amount
		}
	}
	return credits - debits
}

func (s *transactionService) GetIssues() []model.Transaction {
	all := s.repo.GetAll()
	out := make([]model.Transaction, 0)
	for _, t := range all {
		if t.Status == "FAILED" || t.Status == "PENDING" {
			out = append(out, t)
		}
	}
	return out
}
