package main

import (
	"time"

	"github.com/hhq163/breaker"
)

const (
	MSG_NULL_ACTION = 0

	SMSG_SIGNAL_AUTH         = 1
	SMSG_SIGNAL_INFO         = 2
	SMSG_SIGNAL_CLEARHISTORY = 3
	SMSG_SIGNAL_CHANGEDEALER = 4
	SMSG_SIGNAL_LASTGAME     = 5
	SMSG_SIGNAL_MAINTAIN     = 6
	SMSG_SIGNAL_CARDINFO     = 7
	SMSG_CARD_CONNECT        = 8
)

func main() {
	var cmds = []int32{1000}
	var options = breaker.Options{
		BucketTime:        150 * time.Millisecond,
		BucketNums:        6450, //每秒4万次请求，超过这个值熔断
		BreakerRate:       0.6,
		BreakerMinSamples: 300,
		CoolingTimeout:    3 * time.Second,
		DetectTimeout:     150 * time.Millisecond,
		HalfOpenSuccess:   3,
	}
	cb := breaker.InitBreakers(cmds, options)

	if !cb.IsTriggerBreaker(1000) {
		//正常调用
		// if true {
		// 	cb.Succeed()
		// } else {
		// 	cb.Fail()
		// }

		//
	} else {
		return
	}

	// cb.Fail()
	// cb.Timeout()
}
