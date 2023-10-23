package etf

import (
	"time"
)

type ETFData struct {
	Price          float64   // ETF当天价格
	Date           time.Time // 日期
	Symbol         string    // ETF标识
	GDPGrowth      float64   // GDP预估增长率
	EstimatedPrice float64   // 基于GDP增长率的估算价格
	PriceRatio     float64   // 估算价格与实际价格的比率
}

func NewETFData(price float64, date time.Time, symbol string, gdpGrowth float64) *ETFData {
	return &ETFData{
		Price:     price,
		Date:      date,
		Symbol:    symbol,
		GDPGrowth: gdpGrowth,
	}
}

func (e *ETFData) CalculateEstimatedPrice(previousPrice float64, tradingDaysInYear int) {
	growthRate := e.GDPGrowth / float64(tradingDaysInYear)
	e.EstimatedPrice = previousPrice * (1 + growthRate)
}

func (e *ETFData) CalculatePriceRatio() {
	if e.EstimatedPrice != 0 {
		e.PriceRatio = e.Price / e.EstimatedPrice
	}
}
