package authorization

type Authorization uint8

const (
	Token  Authorization = 0
	Bearer Authorization = 1
)

type DiscordAuthorization struct {
	Mode  Authorization
	Token string
}

func (d *DiscordAuthorization) BuildAuthorizationHeader() string {
	switch d.Mode {
	case Token:
		return "Bot " + d.Token
	case Bearer:
		return "Bearer " + d.Token
	default:
		panic("Invalid authorization mode: " + string(d.Mode))
	}
}

func BuildAuthorization(mode Authorization, token string) *DiscordAuthorization {
	return &DiscordAuthorization{
		Mode:  mode,
		Token: token,
	}
}
