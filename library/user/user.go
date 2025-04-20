package user

import (
	"encoding/json"
	"io"

	"disgord.dev/disgord/library/authorization"
	"disgord.dev/disgord/library/rest"
	"disgord.dev/disgord/library/snowflake"
)

const (
	FlagStaff                 = 1 << 0
	FlagPartner               = 1 << 1
	FlagHypesquad             = 1 << 2
	FlagBugHunterLevel1       = 1 << 3
	FlagHouseBravery          = 1 << 6
	FlagHouseBrilliance       = 1 << 7
	FlagHouseBalance          = 1 << 8
	FlagPremiumEarlySupporter = 1 << 9
	FlagTeamUser              = 1 << 10
	FlagBugHunterLevel2       = 1 << 14
	FlagVerifiedBot           = 1 << 16
	FlagVerifiedBotDeveloper  = 1 << 17
	FlagCertifiedModerator    = 1 << 18
	FlagBotHTTPInteractions   = 1 << 19
	FlagActiveDeveloper       = 1 << 22
)

const (
	PremiumNitroClassic = 1
	PremiumNitro        = 2
	PremiumNitroBasic   = 3
)

type DiscordAvatarDecorationData struct {
	Asset string
	SkuID snowflake.Snowflake
}

func ParseDiscordAvatarDecorationData(dataJSON map[string]any) *DiscordAvatarDecorationData {
	return &DiscordAvatarDecorationData{
		Asset: dataJSON["asset"].(string),
		SkuID: snowflake.ParseSnowflake(dataJSON["sku_id"].(uint64)),
	}
}

type DiscordUser struct {
	ID                   snowflake.Snowflake
	Username             string
	Discriminator        string
	GlobalName           string
	Avatar               string
	Bot                  bool
	System               bool
	MFAEnabled           bool
	Banner               string
	AccentColor          int
	Locale               string
	Verified             bool
	Email                string
	Flags                int
	PremiumType          int
	PublicFlags          int
	AvatarDecorationData *DiscordAvatarDecorationData
}

func parseString(data any) string {
	if str, ok := data.(string); ok {
		return str
	}
	return ""
}

func parseBool(data any) bool {
	if b, ok := data.(bool); ok {
		return b
	}
	return false
}

func parseInt(data any) int {
	if i, ok := data.(int); ok {
		return i
	}
	return 0
}

func parseSnowflake(data any) snowflake.Snowflake {
	if str, ok := data.(string); ok {
		return snowflake.ParseSnowflakeString(str)
	}

	return snowflake.Snowflake{}
}

func parseDiscordAvatarDecorationData(data any) *DiscordAvatarDecorationData {
	if dataMap, ok := data.(map[string]any); ok {
		return ParseDiscordAvatarDecorationData(dataMap)
	}
	return nil
}

func ParseDiscordUser(userJSON map[string]any) *DiscordUser {
	return &DiscordUser{
		ID:            parseSnowflake(userJSON["id"]),
		Username:      parseString(userJSON["username"]),
		Discriminator: parseString(userJSON["discriminator"]),
		GlobalName:    parseString(userJSON["global_name"]),
		Avatar:        parseString(userJSON["avatar"]),
		Bot:           parseBool(userJSON["bot"]),
		System:        parseBool(userJSON["system"]),
		MFAEnabled:    parseBool(userJSON["mfa_enabled"]),
		Banner:        parseString(userJSON["banner"]),
		AccentColor:   parseInt(userJSON["accent_color"]),
		Locale:        parseString(userJSON["locale"]),
		Verified:      parseBool(userJSON["verified"]),
		Email:         parseString(userJSON["email"]),
		Flags:         parseInt(userJSON["flags"]),
		PremiumType:   parseInt(userJSON["premium_type"]),
		PublicFlags:   parseInt(userJSON["public_flags"]),

		AvatarDecorationData: parseDiscordAvatarDecorationData(userJSON["avatar_decoration_data"]),
	}
}

func GetCurrentUser(authorization *authorization.DiscordAuthorization) (*DiscordUser, error) {
	response, e := rest.RestRequest(authorization, "GET", rest.GetCurrentUser)

	if e != nil {
		return nil, e
	}

	// if response.StatusCode != 200 {
	// 	return nil, (response.StatusCode, response.Status, response.Body)
	// }

	body, _ := io.ReadAll(response.Body)
	bodyJSON := make(map[string]interface{})

	json.Unmarshal(body, &bodyJSON)

	user := ParseDiscordUser(bodyJSON)

	response.Body.Close()

	return user, nil
}
