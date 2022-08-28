package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once sync.Once
	v    heatTicker
)

type heatTicker struct {
	updateHeat  *time.Ticker     // 更新房间热度的计时器
	updateHeatC <-chan time.Time // channel
	stop        chan bool        // 用于停止所有计时器
}

func init() {
	once.Do(func() {
		var t heatTicker
		t.updateHeat = time.NewTicker(2 * time.Second)
		t.updateHeatC = t.updateHeat.C

		v = t
	})

	fmt.Println("init")
}

func main() {
	fmt.Println("start")

	loop(v)
	fmt.Println("end")
}

func loop(t heatTicker) {
	defer t.updateHeat.Stop()

	for {
		select {
		case <-t.updateHeatC:
			fmt.Println("AAA")
		case <-t.stop:
			return
		}

	}
}
