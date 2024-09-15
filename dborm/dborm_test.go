package dborm

import (
	"testing"

	"github.com/jerryharbour/libgo/dbdriver"
)

type MarketType int

const (
	MarketType_Unknown             = MarketType(0)  // unknow market
	MarketType_HK                  = MarketType(1)  // hongkong market
	MarketType_US                  = MarketType(2)  // usa market
	MarketType_CN                  = MarketType(3)  // china market
	MarketType_HKCC                = MarketType(4)  // A stock in hongkong
	MarketType_Futures             = MarketType(5)  // futures market(global futures)
	MarketType_Futures_Simulate_HK = MarketType(10) // simulate futures market of hongkong
	MarketType_Futures_Simulate_US = MarketType(11) // simulate futures market of usa
	MarketType_Futures_Simulate_SG = MarketType(12) // simulate futures market of singapore
	MarketType_Futures_Simulate_JP = MarketType(13) // simulate futures market of japan
)

type Security struct {
	Market MarketType `json:"market" gorm:"column:market;default:0;not null;"` // market type
	Code   string     `json:"code gorm:"column:code;default:'';not null;"`     // the code of the sock
}

type KLine struct {
	Id             int32   `json:"id" gorm:"column:id;primary_key;autoIncrement;not null"`         // the id of the kline
	Time           string  `json:"time" gorm:"column:time"`                                        // timestamp (format: yyyy-MM-dd HH:mm::ss)
	IsBlank        bool    `json:"isBlank" gorm:"column:isblank;default:false;not null"`           // whether is blank content, if true, only time information
	HighPrice      float64 `json:"highPrice" gorm:"column:high_price"`                             // the highest price
	OpenPrice      float64 `json:"openPrice gorm:"column:open_price"`                              // the opening price
	ClosePrice     float64 `json:"closePrice gorm:"column:close_price"`                            // the closing price
	LowPrice       float64 `json:"lowPrice" gorm:"column:low_price"`                               // the lowest price
	LastClosePrice float64 `json:"lastClosePrice" gorm:"column:yesterday_close_price"`             // the yesterday's closing price
	Volume         int64   `json:"volume" gorm:"column:volume;default:0;not null;"`                // the volume
	Turnover       float64 `json:"turnover" gorm:"column:turnover;default:0.0;not null;"`          // the turnover
	TurnoverRate   float64 `json:"turnoverRate" gorm:"column:turnover_rate;default:0.0;not null;"` // the turnover rate
	PE             float64 `json:"pe" gorm:"column:pe_ratio;default:0.0;not null;"`                // P/E ratio
	ChangeRate     float64 `json:"changeRate" gorm:"column:change_rate;default:0.0;not null;"`     // the change rate
	Timestamp      float64 `json:"timestamp" gorm:"column:timestamp"`                              // the timestamp
}

type KLineData struct {
	KLine
	Security
}

func NewKLineData() *KLineData {
	return &KLineData{
		KLine: KLine{
			Time:           "2024-07-04",
			HighPrice:      100.0,
			OpenPrice:      98.1,
			ClosePrice:     95.2,
			LowPrice:       90.7,
			LastClosePrice: 95.8,
			Volume:         1002000,
			Turnover:       20020.9,
			TurnoverRate:   20.5,
			PE:             0.5,
			ChangeRate:     0.001,
			Timestamp:      1530409600,
		},
		Security: Security{
			Market: MarketType_HK,
			Code:   "00700",
		},
	}
}

func TestMySqlOpen(t *testing.T) {
	mysqlConf := dbdriver.NewMySqlConfig("hk-cdb-3crr0rin.sql.tencentcdb.com", 23572, "root", "mysql-140720", "stock")
	dbOrm := NewDbOrm()
	if err := dbOrm.Open(mysqlConf); err != nil {
		t.Errorf("open mysql error: %v", err)
	}

	if err := dbOrm.InsertIntoTable("kdata_business700", NewKLineData()); err != nil {
		t.Errorf("insert into table error: %v", err)
	}

	dbOrm.Close()
}
