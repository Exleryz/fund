package base

import (
	"fund/conf"
)

type BaseDao struct {
	DataSource conf.DataSource
}

func (baseDao *BaseDao) GetDataSource() conf.DataSource {
	return baseDao.DataSource
}
