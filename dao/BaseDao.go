package dao

import (
	"fund/conf"
	"fund/dao/xq"
	"github.com/shopspring/decimal"
)

type BaseDao struct {
	DataSource conf.DataSource
}

type FundDao interface {
	GetFundPrice(fundCode string) decimal.Decimal
	GetDataSource() conf.DataSource
}

func NewFundDao(dataSource conf.DataSource) FundDao {
	switch dataSource {
	case conf.XueQiu:
		// 因为 是指针类型
		return &xq.XueQiuFundDao{}
	case conf.JiSiLu:
	case conf.SINA:
	default:
		return nil
	}
	return nil
}
