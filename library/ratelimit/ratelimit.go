package ratelimit

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RateLimit struct {
	ResetEpoch int
	RWMutex    *sync.RWMutex
}

var bucketLimits = map[string]*RateLimit{}

func validateBucketMap(bucket string) {
	if _, ok := bucketLimits[bucket]; !ok {
		bucketLimits[bucket] = &RateLimit{
			ResetEpoch: 0,
			RWMutex:    &sync.RWMutex{},
		}
	}
}

func RateLimitRequest(bucket string) {
	validateBucketMap(bucket)

	bucketLimits[bucket].RWMutex.RLock()

	currentEpoch := int(time.Now().Unix())
	timeUntilReset := bucketLimits[bucket].ResetEpoch - currentEpoch

	fmt.Println("RateLimitRequest: ", bucket, " - ", timeUntilReset)

	if timeUntilReset > 0 {
		bucketLimits[bucket].RWMutex.RUnlock()

		time.Sleep(time.Duration(timeUntilReset) * time.Second)
	} else {
		bucketLimits[bucket].RWMutex.RUnlock()
	}
}

func ParseRateLimitBucket(header *http.Header) string {
	return header.Get("X-RateLimit-Bucket")
}

func ParseRateLimitHeader(header *http.Header) {
	validateBucketMap(ParseRateLimitBucket(header))

	if reset := header.Get("X-RateLimit-Reset"); reset != "" {
		bucket := ParseRateLimitBucket(header)

		bucketLimits[bucket].RWMutex.Lock()

		ratelimitParsed, _ := strconv.ParseInt(header.Get("X-RateLimit-Reset"), 10, 64)

		bucketLimits[bucket].ResetEpoch = int(time.Now().Unix()) + int(ratelimitParsed)

		bucketLimits[bucket].RWMutex.Unlock()
	}

}
