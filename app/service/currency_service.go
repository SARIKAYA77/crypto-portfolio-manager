package service

import (
	"crypto-portfolio/app/model"
	"crypto-portfolio/app/util"
	_ "crypto-portfolio/app/util"
	"errors"
	"log"
)

type CurrencyService struct {
	currencies []*model.Currency
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		currencies: make([]*model.Currency, 0),
	}
}

func (s *CurrencyService) AddCurrency(code string, amount float64) (*model.Currency, error) {
	for _, c := range s.currencies {
		if c.Code == code {
			return nil, errors.New("currency already exists")
		}
	}

	price, err := util.GetCryptoPrice(code)
	if err != nil {
		return nil, err
	}
	totalValue := price * amount
	newCurrency := &model.Currency{
		ID:         len(s.currencies) + 1,
		Code:       code,
		Amount:     amount,
		Price:      price,
		TotalValue: totalValue,
		History:    []model.Price{},
	}
	s.currencies = append(s.currencies, newCurrency)
	log.Printf("Added currency with code: %s, amount: %f, total value: %f", code, amount, totalValue)
	return newCurrency, nil
}

func (s *CurrencyService) EditCurrency(id int, code string, amount float64) (*model.Currency, error) {
	for _, c := range s.currencies {
		if c.ID == id {
			price, err := util.GetCryptoPrice(code)
			if err != nil {
				return nil, err
			}
			totalValue := price * amount
			c.Code = code
			c.Amount = amount
			c.Price = price
			c.TotalValue = totalValue
			log.Printf("Edited currency with ID: %d, new code: %s, new amount: %f, new total value: %f", id, code, amount, totalValue)
			return c, nil
		}
	}

	return nil, errors.New("currency with that ID does not exist")
}

func (s *CurrencyService) DeleteCurrency(id int) error {
	for i, c := range s.currencies {
		if c.ID == id {
			s.currencies = append(s.currencies[:i], s.currencies[i+1:]...)

			log.Printf("Deleted currency with ID: %d", id)

			return nil
		}
	}

	return errors.New("currency with that ID does not exist")
}

func (s *CurrencyService) GetCurrencyByID(id int) (*model.Currency, error) {
	for _, c := range s.currencies {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, errors.New("currency with that ID does not exist")
}

func (s *CurrencyService) GetAllCurrencies() ([]*model.Currency, error) {
	return s.currencies, nil
}
