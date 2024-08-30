package mc_server

import (
	fb_client "Eulogist/core/fb_auth/mv4/client"
	"Eulogist/core/minecraft/protocol/login"
	raknet_connection "Eulogist/core/raknet"
	"Eulogist/core/tools/skin_process"
)

type MinecraftServer struct {
	fbClient              *fb_client.Client
	authResponse          *fb_client.AuthResponse
	getCheckNumEverPassed bool

	identityData *login.IdentityData
	clientData   *login.ClientData

	neteaseUID string
	playerSkin *skin_process.Skin
	outfitInfo map[string]*int

	entityUniqueID  int64
	entityRuntimeID uint64

	*raknet_connection.Raknet
}

// ...
type BasicConfig struct {
	ServerCode     string
	ServerPassword string
	Token          string
	AuthServer     string
}
