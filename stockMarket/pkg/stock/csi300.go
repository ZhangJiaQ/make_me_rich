package etf

import "time"

// CSI300  用于表示沪深300指数的数据结构。
type CSI300 struct {
	DirectionColor   string    `json:"direction_color"`   // 方向颜色，用于前端显示不同的文本颜色
	RowDate          string    `json:"rowDate"`           // 日期的字符串表示，格式为"2008年12月30日"
	RowDateRaw       int64     `json:"rowDateRaw"`        // 日期的原始整数表示
	RowDateTimestamp time.Time `json:"rowDateTimestamp"`  // 日期的时间戳表示，使用ISO 8601格式
	LastClose        string    `json:"last_close"`        // 上一个交易日的收盘价，字符串表示
	LastOpen         string    `json:"last_open"`         // 当天的开盘价，字符串表示
	LastMax          string    `json:"last_max"`          // 当天的最高价，字符串表示
	LastMin          string    `json:"last_min"`          // 当天的最低价，字符串表示
	Volume           string    `json:"volume"`            // 成交量，字符串表示
	VolumeRaw        int       `json:"volumeRaw"`         // 成交量，整数表示
	ChangePercent    string    `json:"change_precent"`    // 变化百分比，字符串表示
	LastCloseRaw     float64   `json:"last_closeRaw"`     // 上一个交易日的收盘价，浮点数表示
	LastOpenRaw      float64   `json:"last_openRaw"`      // 当天的开盘价，浮点数表示
	LastMaxRaw       float64   `json:"last_maxRaw"`       // 当天的最高价，浮点数表示
	LastMinRaw       float64   `json:"last_minRaw"`       // 当天的最低价，浮点数表示
	ChangePercentRaw float64   `json:"change_precentRaw"` // 变化百分比，浮点数表示
}
