package rest

import (
	"fmt"
	"net/http"

	"disgord.dev/disgord/library/authorization"
	"disgord.dev/disgord/library/ratelimit"
)

func buildRestHeaders(auth *authorization.DiscordAuthorization) map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "DiscordBot (disgord.dev, 0.0.0)",

		"Authorization": auth.BuildAuthorizationHeader(),
	}
}

func buildRestRequest(authorization *authorization.DiscordAuthorization, method string, endpoint string) (*http.Request, error) {
	request, err := http.NewRequest(method, fmt.Sprintf("https://discord.com/api/v10/%s", endpoint), nil)

	if err != nil {
		return nil, err
	}

	headers := buildRestHeaders(authorization)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	return request, nil
}

func RestRequest(authorization *authorization.DiscordAuthorization, method string, endpoint string) (*http.Response, error) {
	request, err := buildRestRequest(authorization, method, endpoint)

	if err != nil {
		return nil, err
	}

	ratelimit.ParseRateLimitHeader(&request.Header)

	bucket := ratelimit.ParseRateLimitBucket(&request.Header)

	SendChannelRequest(bucket, request)

	response := GetChannelRequest(bucket)

	if response.Response != nil {
		ratelimit.ParseRateLimitHeader(&response.Response.Header)

		return response.Response, nil
	}

	return nil, fmt.Errorf("failed to get response from channel: %s", bucket)
}
