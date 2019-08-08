package times

import (
	"math/rand"
	"time"
)

func RandomDurationSeconds(start, end int64) time.Duration {
	rand.Seed(time.Now().UnixNano())

	return time.Duration(rand.Int63n(end-start)+start) * time.Second
}

func RandomSeconds(duration int64) time.Duration {
	return RandomDurationSeconds(1, duration)
}
