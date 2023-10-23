package scraper

import (
	"github.com/gocolly/colly/v2"
	"time"
)

type ETFCrawler struct {
	colly *colly.Collector
}

type ETFPriceData struct {
	Date  time.Time
	Price float64
}

func NewETFCrawler() *ETFCrawler {
	c := colly.NewCollector()

	return &ETFCrawler{colly: c}
}

func (ec *ETFCrawler) FetchETFPricing(url string) ([]ETFPriceData, error) {
	var data []ETFPriceData
	ec.colly.OnHTML("your-css-selector-for-price-and-date", func(e *colly.HTMLElement) {
		// 解析日期和价格数据，并添加到data切片中
		// 你可能需要使用strconv库将字符串转换为float64和time包解析日期
	})

	ec.colly.OnError(func(r *colly.Response, err error) {
		// handle error
	})

	err := ec.colly.Visit(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
