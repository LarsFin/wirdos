package util

import "time"

var last = time.Now()

var DeltaTime = float64(0)

func UpdateDeltaTime() {
	DeltaTime = time.Since(last).Seconds()
	last = time.Now()
}
