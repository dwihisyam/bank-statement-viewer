package repository

import (
	"sync"

	"bank-statement-viewer-backend/internal/model"
)

type TransactionRepository interface {
	SaveAll([]model.Transaction) error
	AppendAll([]model.Transaction) error
	GetAll() []model.Transaction
	Clear() // optional helper
}

type inMemoryRepo struct {
	mu   sync.RWMutex
	data []model.Transaction
}

func NewInMemoryRepo() TransactionRepository {
	return &inMemoryRepo{
		data: make([]model.Transaction, 0),
	}
}

func (r *inMemoryRepo) SaveAll(tx []model.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = make([]model.Transaction, 0, len(tx))
	r.data = append(r.data, tx...)
	return nil
}

func (r *inMemoryRepo) AppendAll(tx []model.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = append(r.data, tx...)
	return nil
}

func (r *inMemoryRepo) GetAll() []model.Transaction {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]model.Transaction, len(r.data))
	copy(out, r.data)
	return out
}

func (r *inMemoryRepo) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = r.data[:0]
}
