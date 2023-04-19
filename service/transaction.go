package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"app/config"
	"app/model"
)

type transactionService struct {
	baseURL string
	client  *http.Client
}

func NewTransactionService(conf *config.Config) *transactionService {
	return &transactionService{
		baseURL: conf.Transaction.BaseURL,
		client:  &http.Client{},
	}
}

const transactionServicePath = `transactionServise: %w`

func (s *transactionService) Create(ctx context.Context, tr model.CreateTransactionRq) (rp model.CreateTransactionRp, err error) {
	code, err := s.doReq(ctx, http.MethodPost, "/transaction", tr, rp)
	if err != nil {
		return rp, fmt.Errorf(transactionServicePath, err)
	}

	if code != http.StatusCreated {
		return rp, fmt.Errorf(transactionServicePath, model.ErrTransactionFailed)
	}

	return rp, nil
}

func (s *transactionService) Pay(ctx context.Context, tr model.PayTransactionRq) error {
	code, err := s.doReq(ctx, http.MethodPatch, "/transaction", tr, nil)
	if err != nil {
		return fmt.Errorf(transactionServicePath, err)
	}

	if code != http.StatusOK {
		return fmt.Errorf(transactionServicePath, model.ErrTransactionFailed)
	}

	return nil
}

func (s *transactionService) doReq(ctx context.Context, method, source string, input, dest interface{}) (status int, err error) {
	data, err := json.Marshal(input)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(method, s.baseURL+source, bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	req = req.WithContext(ctx)

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if dest != nil {
		err = json.NewDecoder(resp.Body).Decode(dest)
		if err != nil {
			return 0, err
		}
	}

	return resp.StatusCode, nil
}
