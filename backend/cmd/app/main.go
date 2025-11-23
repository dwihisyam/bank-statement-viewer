package main

import (
	"log"
	"net/http"
	"time"

	"bank-statement-viewer-backend/internal/handler"
	"bank-statement-viewer-backend/internal/repository"
	"bank-statement-viewer-backend/internal/service"
)

func main() {
	repo := repository.NewInMemoryRepo()
	svc := service.NewTransactionService(repo)
	h := handler.NewTransactionHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/upload", h.UploadHandler)
	mux.HandleFunc("/balance", h.BalanceHandler)
	mux.HandleFunc("/issues", h.IssuesHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      corsMiddleware(loggingMiddleware(mux)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("Server running on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		_ = start // could log if wanted
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request directly
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
