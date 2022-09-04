package utils

import (
	"fmt"
	"runtime"
	"time"
)

func TimeCost() func() {
	startTime := time.Now()
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return func() {
		t := time.Since(startTime)
		fmt.Println(f.Name(), " time cost : ", t.Nanoseconds())
	}
}
