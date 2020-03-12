package indicators

import (
	"fmt"

	objects "github.com/d5/tengo/v2"
)

// RsiModule relative strength index indicator commands
var RsiModule = map[string]objects.Object{
	"rsi": &objects.UserFunction{Name: "rsi", Value: rsi},
}

func rsi(args ...objects.Object) (objects.Object, error) {
	if len(args) != 2 {
		return nil, objects.ErrWrongNumArguments
	}
	ohlcData := objects.ToInterface(args[0])
	x := ohlcData.(map[string]interface{})

	total := len(x)
	//fmt.Println(x["timestamp"])
	//
	// ohlcData := objects.ToInterface(args[0])
	// ohlcTimestampData, err := appendDataInt64(ohlcData.([]interface{}))
	// if err != nil {
	// 	return nil, err
	// }
	//
	// ohlcData = objects.ToInterface(args[0])
	// ohlcCloseData, err := appendDataFloat(ohlcData.([]interface{}))
	// if err != nil {
	// 	return nil, err
	// }
	//
	// inTimePeroid, ok := objects.ToInt(args[1])
	// if !ok {
	// 	return nil, fmt.Errorf(modules.ErrParameterConvertFailed, inTimePeroid)
	// }
	//
	// for x := range ohlcTimestampData
	// period := techan.NewTimePeriod(time.Unix(ohlcTimestampData[0], 0), time.Hour*24)
	//
	// candle := techan.NewCandle(period)
	// candle.ClosePrice = big.NewFloat()
	// ret := indicators.Rsi(ohlcCloseData, inTimePeroid)
	r := &objects.Array{}
	// for x := range ret {
	// 	r.Value = append(r.Value, &objects.Float{Value: ret[x]})
	// }

	return r, nil
}
