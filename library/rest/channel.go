package rest

import (
	"net/http"

	"disgord.dev/disgord/library/ratelimit"
)

type ChannelCommunication struct {
	Request  *http.Request
	Response *http.Response
}

var bucketChannels = map[string]chan *ChannelCommunication{}

func validateBucketMap(bucket string) {
	if _, ok := bucketChannels[bucket]; !ok {
		bucketChannels[bucket] = make(chan *ChannelCommunication, 1)
	}
}

func SendChannelRequest(bucket string, request *http.Request) {
	validateBucketMap(bucket)

	go func() {
		ratelimit.RateLimitRequest(bucket)

		response, err := http.DefaultClient.Do(request)

		if err != nil {
			return
		}

		bucketChannels[bucket] <- &ChannelCommunication{
			Response: response,
		}
	}()
}

func GetChannelRequest(bucket string) *ChannelCommunication {
	validateBucketMap(bucket)

	return <-bucketChannels[bucket]
}
