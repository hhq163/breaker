# breaker

breaker is a simple circuitBreaker, you can use it directly, or you can read and reform it .
base on github.com/JeffreyDing11223/goBreaker

## easy to start:

```go
go get github.com/hhq163/breaker
```

## easy to use:
```go

  var cmds = []int32{1, 2, 289, 55}
	var options = breaker.Options{
		BucketTime:        150 * time.Millisecond,
		BucketNums:        200,
		BreakerRate:       0.6,
		BreakerMinSamples: 300,
		CoolingTimeout:    3 * time.Second,
		DetectTimeout:     150 * time.Millisecond,
		HalfOpenSuccess:   3,
	}
	cb := breaker.InitCircuitBreakers(cmds, options)
	
  ...
  
	if cb.IsTriggerBreaker(289){
		// downStream service is broken
	}
  
```
