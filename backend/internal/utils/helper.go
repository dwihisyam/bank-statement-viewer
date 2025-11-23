package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"bank-statement-viewer-backend/internal/model"
)

func ParseCSVTransactions(r io.Reader) ([]model.Transaction, error) {
	cr := csv.NewReader(r)
	cr.TrimLeadingSpace = true
	cr.FieldsPerRecord = -1 // allow variable
	transactions := make([]model.Transaction, 0)

	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("csv read error: %w", err)
		}
		// skip empty lines
		if len(record) == 0 {
			continue
		}
		// Expect at least 6 columns, but allow extra columns
		if len(record) < 6 {
			return nil, fmt.Errorf("invalid record length: %v", record)
		}

		tsStr := strings.TrimSpace(record[0])
		name := strings.TrimSpace(record[1])
		typ := strings.ToUpper(strings.TrimSpace(record[2]))
		amtStr := strings.TrimSpace(record[3])
		status := strings.ToUpper(strings.TrimSpace(record[4]))
		desc := strings.TrimSpace(record[5])

		ts, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp '%s': %w", tsStr, err)
		}
		amt, err := strconv.ParseInt(amtStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid amount '%s': %w", amtStr, err)
		}

		t := model.Transaction{
			Timestamp:   ts,
			Name:        name,
			Type:        typ,
			Amount:      amt,
			Status:      status,
			Description: desc,
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
