package discordgo

import (
	"disgord.dev/disgord/library/authorization"
	"disgord.dev/disgord/library/user"
)

type DiscordApplication struct {
	Authorization *authorization.DiscordAuthorization
}

func BuildDiscordApplicationToken(token string) *DiscordApplication {
	return &DiscordApplication{
		Authorization: authorization.BuildAuthorization(authorization.Token, token),
	}
}

func (application *DiscordApplication) GetUser() (*user.DiscordUser, error) {
	return user.GetCurrentUser(application.Authorization)
}
