package enums

type PlatformMenuType int

const (
	MenuPriPlatform PlatformMenuType = 1 //只针对平台
	MenuPub         PlatformMenuType = 2 //共享
	MenuPriTeam     PlatformMenuType = 3 //团队
)
