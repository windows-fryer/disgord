package rest

import (
	"fmt"
	"net/http"

	"disgord.dev/disgord/internal/authorization"
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

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
