package discordgo

import (
	"encoding/json"
	"io"

	"disgord.dev/disgord/internal/authorization"
	"disgord.dev/disgord/internal/rest"
	"disgord.dev/disgord/internal/user"
)

type DiscordApplication struct {
	Authorization *authorization.DiscordAuthorization
}

func BuildDiscordApplicationToken(token string) *DiscordApplication {
	return &DiscordApplication{
		Authorization: authorization.BuildAuthorization(authorization.Token, token),
	}
}

func (application *DiscordApplication) IdentifyApplication() (*user.DiscordUser, error) {
	response, e := rest.RestRequest(application.Authorization, "GET", rest.GetCurrentUser)

	if e != nil {
		return nil, e
	}

	body, _ := io.ReadAll(response.Body)
	bodyJSON := make(map[string]interface{})

	json.Unmarshal(body, &bodyJSON)

	user := user.ParseDiscordUser(bodyJSON)

	response.Body.Close()

	return user, nil
}
