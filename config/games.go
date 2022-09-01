package config

var (
	DungeonCooldown = getEnvInt("DUNGEON_COOLDOWN")
	GambleCooldown  = getEnvInt("GAMBLE_COOLDOWN")
)
