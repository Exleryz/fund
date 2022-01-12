# 使用方法

**持仓不构成投资建议**

## Windows 使用cmd执行

```shell
StockHis.exe -sc SH601318 -p day
```

参数详解

- sc: 股票代码
    - A股:SH/SZ + 股票代码
    - H股:09988
    - M股:TSLA
- p: 数据周期
    - 1m: 一分钟
    - 5m
    - 15m
    - 60m
    - 120m
    - day: 一天
    - week: 一周
    - quarter
    - year: 一年

执行完命令后会在当前文件夹下生成一个`tmp\<sc-p>.csv`文件

## Mac / Linux

# csv文件字段详情

| timestamp | volume | open   | high   | low    | close  | chg    | percent | turnoverrate | amount | volume_post | amount_post | pe  | pb  | ps  | pcf | market_capital | balance | hold_volume_cn | hold_ratio_cn | net_volume_cn | hold_volume_hk | hold_ratio_hk | net_volume_hk |
| --------- | ------ | ------ | ------ | ------ | ------ | ------ | ------- | ------------ | ------ | ----------- | ----------- | --- | --- | --- | --- | -------------- | ------- | -------------- | ------------- | ------------- | -------------- | ------------- | ------------- |
| 日期      | 成交量 | 开盘价 | 最高价 | 最低价 | 收盘价 | 涨跌额 | 涨跌幅  | 换手率       | 成交额 |             |             |     |     |     |     |                |         |                |               |               |                |               |               |

# 数据来源

1. 雪球

# calc

用于计算 calc/conf/position.yml 文件中配置的持仓股票/基金

```yaml
StockList:
  - stock:
      # 名称
      name: 招商银行
      # 必填 股票/基金编码
      code: SH600036
      # 必填 0: 股票 1: 基金
      type: 0
      # 必填 持仓数量
      count: 100
      # 购买平台
      source: 3
```