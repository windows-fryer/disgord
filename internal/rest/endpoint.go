package rest

const (
	GetCurrentUser                          = "users/@me"
	GetUser                                 = "users/%s"
	ModifyCurrentUser                       = "users/@me"
	GetCurrentUserGuilds                    = "users/@me/guilds"
	GetCurrentUserGuildMember               = "users/@me/guilds/%s/member"
	LeaveGuild                              = "users/@me/guilds/%s"
	CreateDM                                = "users/@me/channels"
	CreateGroupDM                           = "users/@me/channels"
	GetCurrentUserConnections               = "users/@me/connections"
	GetCurrentuserApplicatonRoleConnections = "users/@me/applications/%s/role-connections"
)
