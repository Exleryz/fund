package service

import "github.com/shopspring/decimal"

type AssetsService interface {
	GetPrice(assetsCode string) decimal.Decimal
}
