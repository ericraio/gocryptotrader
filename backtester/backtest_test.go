package backtest

import (
	"fmt"
	"testing"
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/kline"
	gctorder "github.com/thrasher-corp/gocryptotrader/exchanges/order"
	"golang.org/x/exp/rand"
)

type testBT struct{}

func (bt *testBT) Init() *Config {
	return &Config{
		Item: genOHCLVData(),
	}
}

func (bt *testBT) OnData(d DataEvent, b *Backtest) (bool, error) {
	b.Portfolio.Order(1.2, 1, gctorder.Sell)
	return true, nil
}

func (bt *testBT) OnEnd(b *Backtest) {
	// b.Stats.PrintResult()
	fmt.Println(b.Stats.ReturnResult())
}

func TestBacktest_Run(t *testing.T) {
	g := &testBT{}
	Run(g)
}

func genOHCLVData() (outItem kline.Item) {
	start := time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

	outItem.Interval = kline.OneDay
	outItem.Asset = asset.Spot
	outItem.Pair = currency.NewPair(currency.BTC, currency.USDT)
	outItem.Exchange = "test"

	outItem.Candles = append(outItem.Candles, kline.Candle{
		Time:   start,
		Open:   0,
		High:   10 + rand.Float64(),
		Low:    10 + rand.Float64(),
		Close:  10 + rand.Float64(),
		Volume: 10,
	})

	for x := 1; x < 365; x++ {
		outItem.Candles = append(outItem.Candles, kline.Candle{
			Time:   start.Add(time.Hour * 24 * time.Duration(x)),
			Open:   outItem.Candles[x-1].Close,
			High:   outItem.Candles[x-1].Open + rand.Float64(),
			Low:    outItem.Candles[x-1].Open - rand.Float64(),
			Close:  outItem.Candles[x-1].Open + rand.Float64(),
			Volume: 10,
		})
	}

	return outItem
}
