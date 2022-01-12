package service

import (
	"fund/conf"
	"fund/dao"
	"github.com/shopspring/decimal"
)

type FundService struct {
	// 数据源
	fundDao dao.FundDao
}

func NewFundService(dataSource conf.DataSource) *FundService {
	return &FundService{fundDao: dao.NewFundDao(dataSource)}
}

func (fundService *FundService) GetPrice(fundCode string) decimal.Decimal {
	return fundService.fundDao.GetFundPrice(fundCode)
}
