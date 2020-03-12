package indicators

import (
	"fmt"

	objects "github.com/d5/tengo/v2"
	"github.com/thrasher-corp/go-talib/indicators"
	"github.com/thrasher-corp/gocryptotrader/gctscript/modules"
)

// MovingAverageModule moving average indicator commands
var MovingAverageModule = map[string]objects.Object{
	"macd": &objects.UserFunction{Name: "macd", Value: macd},
	"ema":  &objects.UserFunction{Name: "ema", Value: ema},
	"sma":  &objects.UserFunction{Name: "sma", Value: sma},
}

func macd(args ...objects.Object) (objects.Object, error) {
	if len(args) != 4 {
		return nil, objects.ErrWrongNumArguments
	}

	ohlcData := objects.ToInterface(args[0])
	ohlcCloseData, err := appendDataFloat(ohlcData.([]interface{}))
	if err != nil {
		return nil, err
	}

	inFastPeriod, ok := objects.ToInt(args[1])
	if !ok {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, inFastPeriod)
	}
	inSlowPeriod, ok := objects.ToInt(args[2])
	if !ok {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, inSlowPeriod)
	}
	inTimePeroid, ok := objects.ToInt(args[3])
	if !ok {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, inTimePeroid)
	}

	macd, macdSignal, macdHist := indicators.Macd(ohlcCloseData, inFastPeriod, inSlowPeriod, inTimePeroid)

	retMACD := &objects.Array{}
	for x := range macd {
		retMACD.Value = append(retMACD.Value, &objects.Float{Value: macd[x]})
	}

	retMACDSignal := &objects.Array{}
	for x := range macdSignal {
		retMACDSignal.Value = append(retMACDSignal.Value, &objects.Float{Value: macdSignal[x]})
	}

	retMACDHist := &objects.Array{}
	for x := range macdHist {
		retMACDHist.Value = append(retMACDHist.Value, &objects.Float{Value: macdHist[x]})
	}

	ret := &objects.Array{}
	ret.Value = append(ret.Value, retMACD, retMACDSignal, retMACDSignal)
	return ret, nil
}

func ema(args ...objects.Object) (objects.Object, error) {
	if len(args) != 2 {
		return nil, objects.ErrWrongNumArguments
	}

	ohlcData := objects.ToInterface(args[0])
	ohlcCloseData, err := appendDataFloat(ohlcData.([]interface{}))
	if err != nil {
		return nil, err
	}

	inTimePeroid, ok := objects.ToInt(args[1])
	if !ok {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, inTimePeroid)
	}

	ret := indicators.Ma(ohlcCloseData, inTimePeroid, indicators.EMA)
	r := &objects.Array{}
	for x := range ret {
		r.Value = append(r.Value, &objects.Float{Value: ret[x]})
	}

	return r, nil
}

func sma(args ...objects.Object) (objects.Object, error) {
	if len(args) != 2 {
		return nil, objects.ErrWrongNumArguments
	}

	ohlcData := objects.ToInterface(args[0])
	ohlcCloseData, err := appendDataFloat(ohlcData.([]interface{}))
	if err != nil {
		return nil, err
	}

	inTimePeroid, ok := objects.ToInt(args[1])
	if !ok {
		return nil, fmt.Errorf(modules.ErrParameterConvertFailed, inTimePeroid)
	}

	ret := indicators.Ma(ohlcCloseData, inTimePeroid, indicators.SMA)
	r := &objects.Array{}
	for x := range ret {
		r.Value = append(r.Value, &objects.Float{Value: ret[x]})
	}

	return r, nil
}
