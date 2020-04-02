# breaker

breaker is a simple circuitBreaker, you can use it directly, or you can read and reform it .
base on github.com/JeffreyDing11223/goBreaker


## example:
```go

 options := breaker.Options{
		BucketTime:        150 * time.Millisecond,
		BucketNums:        6450, //每秒4万次请求，超过这个值熔断
		BreakerRate:       0.3,  //错误率阀值
		BreakerMinSamples: 300,
		CoolingTimeout:    3 * time.Second,        //冷却时间，打开后，过冷却时间后变成半打开
		DetectTimeout:     150 * time.Millisecond, //检测时间，半打开状态以检测时间去发送请求，成功次数到达HalfOpenSuccess后，关闭熔断器
		HalfOpenSuccess:   3,
	}
	breakers := breaker.InitBreakers([]int32{1000}, options)
	cpBreaker = breakers.GetBreaker(1000)

	if cpBreaker.IsAllowed() { //是否被熔断
		// err := actiondo()
		// if err != nil {
		// 	cpBreaker.Fail()
		// 	base.Log.Error("Failed to publish a message i:", i, "errinfo:", err.Error())
		// } else {
		// 	cpBreaker.Succeed()
		// }
	}
  
```
