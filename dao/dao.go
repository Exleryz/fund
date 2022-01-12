package dao

import (
	"fund/conf"
	"fund/dao/base"
	"fund/dao/xq"
	"github.com/shopspring/decimal"
)

type FundDao interface {
	GetFundPrice(fundCode string) decimal.Decimal
	GetDataSource() conf.DataSource
}

func NewFundDao(dataSource conf.DataSource) FundDao {
	switch dataSource {
	case conf.XueQiu:
		// 因为 是指针类型
		return &xq.XueQiuFundDao{BaseDao: base.BaseDao{DataSource: dataSource}}
	case conf.JiSiLu:
	case conf.SINA:
	default:
		return nil
	}
	return nil
}

type StockDao interface {
	// GetStockPrice 返回收盘价
	GetStockPrice(stockCode string) decimal.Decimal
	GetDataSource() conf.DataSource
}

func NewStockDao(dataSource conf.DataSource, period string) StockDao {
	switch dataSource {
	case conf.XueQiu:
		// 因为 是指针类型
		return &xq.XueQiuStockDao{BaseDao: base.BaseDao{DataSource: dataSource}, Period: period}
	case conf.JiSiLu:
	case conf.SINA:
	default:
		return nil
	}
	return nil
}
