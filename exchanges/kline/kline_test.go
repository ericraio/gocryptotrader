package kline

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/thrasher-corp/gocryptotrader/common/crypto"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

func TestValidateData(t *testing.T) {
	err := validateData(nil)
	if err == nil {
		t.Error("error cannot be nil")
	}

	var empty []order.TradeHistory
	err = validateData(empty)
	if err == nil {
		t.Error("error cannot be nil")
	}

	tn := time.Now()
	trade1 := []order.TradeHistory{
		{Timestamp: tn.Add(2 * time.Minute), TID: "2"},
		{Timestamp: tn.Add(time.Minute), TID: "1"},
		{Timestamp: tn.Add(3 * time.Minute), TID: "3"},
	}

	err = validateData(trade1)
	if err == nil {
		t.Error("error cannot be nil")
	}

	trade2 := []order.TradeHistory{
		{Timestamp: tn.Add(2 * time.Minute), TID: "2", Amount: 1, Price: 0},
	}

	err = validateData(trade2)
	if err == nil {
		t.Error("error cannot be nil")
	}

	trade3 := []order.TradeHistory{
		{TID: "2", Amount: 1, Price: 0},
	}

	err = validateData(trade3)
	if err == nil {
		t.Error("error cannot be nil")
	}

	trade4 := []order.TradeHistory{
		{Timestamp: tn.Add(2 * time.Minute), TID: "2", Amount: 1, Price: 1000},
		{Timestamp: tn.Add(time.Minute), TID: "1", Amount: 1, Price: 1001},
		{Timestamp: tn.Add(3 * time.Minute), TID: "3", Amount: 1, Price: 1001.5},
	}

	err = validateData(trade4)
	if err != nil {
		t.Error(err)
	}

	if trade4[0].TID != "1" || trade4[1].TID != "2" || trade4[2].TID != "3" {
		t.Error("trade history sorted incorrectly")
	}
}

func TestCreateKline(t *testing.T) {
	c, err := CreateKline(nil,
		OneMin,
		currency.NewPair(currency.BTC, currency.USD),
		asset.Spot,
		"Binance")
	if err == nil {
		t.Fatal("error cannot be nil")
	}

	var trades []order.TradeHistory
	rand.Seed(time.Now().Unix())
	for i := 0; i < 24000; i++ {
		trades = append(trades, order.TradeHistory{
			Timestamp: time.Now().Add((time.Duration(rand.Intn(10)) * time.Minute) +
				(time.Duration(rand.Intn(10)) * time.Second)),
			TID:    crypto.HexEncodeToString([]byte(string(i))),
			Amount: float64(rand.Intn(20)) + 1,
			Price:  1000 + float64(rand.Intn(1000)),
		})
	}

	c, err = CreateKline(trades,
		0,
		currency.NewPair(currency.BTC, currency.USD),
		asset.Spot,
		"Binance")
	if err == nil {
		t.Fatal("error cannot be nil")
	}

	c, err = CreateKline(trades,
		OneMin,
		currency.NewPair(currency.BTC, currency.USD),
		asset.Spot,
		"Binance")
	if err != nil {
		t.Fatal(err)
	}

	if len(c.Candles) == 0 {
		t.Fatal("no data returned, expecting a lot.")
	}
}

func TestKlineWord(t *testing.T) {
	if OneDay.Word() != "oneday" {
		t.Fatalf("unexpected result: %v", OneDay.Word())
	}
}

func TestKlineDuration(t *testing.T) {
	if OneDay.Duration() != time.Hour*24 {
		t.Fatalf("unexpected result: %v", OneDay.Duration())
	}
}

func TestKlineShort(t *testing.T) {
	if OneDay.Short() != "24h" {
		t.Fatalf("unexpected result: %v", OneDay.Short())
	}
}

func TestDurationToWord(t *testing.T) {
	testCases := []struct {
		name     string
		interval Interval
	}{
		{
			"OneMin",
			OneMin,
		},
		{
			"ThreeMin",
			ThreeMin,
		},
		{
			"FiveMin",
			FiveMin,
		},
		{
			"TenMin",
			TenMin,
		},
		{
			"FifteenMin",
			FifteenMin,
		},
		{
			"ThirtyMin",
			ThirtyMin,
		},
		{
			"OneHour",
			OneHour,
		},
		{
			"TwoHour",
			TwoHour,
		},
		{
			"FourHour",
			FourHour,
		},
		{
			"SixHour",
			SixHour,
		},
		{
			"EightHour",
			OneHour * 8,
		},
		{
			"TwelveHour",
			TwelveHour,
		},
		{
			"OneDay",
			OneDay,
		},
		{
			"ThreeDay",
			ThreeDay,
		},
		{
			"FifteenDay",
			Fifteenday,
		},
		{
			"OneWeek",
			OneWeek,
		},
		{
			"TwoWeek",
			TwoWeek,
		},
		{
			"notfound",
			Interval(time.Hour * 1337),
		},
	}
	for x := range testCases {
		test := testCases[x]
		t.Run(test.name, func(t *testing.T) {
			v := DurationToWord(test.interval)
			if !strings.EqualFold(v, test.name) {
				t.Fatalf("%v: received %v expected %v", test.name, v, test.name)
			}
		})
	}
}

func TestKlineErrors(t *testing.T) {
	v := ErrorKline{
		Interval: OneYear,
	}

	if v.Error() != "oneyear interval unsupported by exchange" {
		t.Fatal("unexpected error returned")
	}

	if v.Unwrap().Error() != "8760h0m0s interval unsupported by exchange" {
		t.Fatal("unexpected error returned")
	}
}

func TestTotalCandlesPerInterval(t *testing.T) {
	end := time.Now()
	start := end.AddDate(-1,0,0)

	v := TotalCandlesPerInterval(start, end , OneYear)
	t.Log(v)

	v = TotalCandlesPerInterval(end.Add(-1 * time.Hour), end, Fifteenday)
	t.Log(v)
}

func TestCalcDateRanges(t *testing.T) {
	end := time.Now()
	start := end.AddDate(-1,0,0)

	v := CalcDateRanges(start, end, OneDay, 300)
	for x := range v {
		for y := range v[x] {
			fmt.Println(v[x][y].start, v[x][y].end)
		}
	}
}