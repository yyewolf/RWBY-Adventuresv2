package config

import (
	"os"
	chars "rwby-adventures/characters"
	"rwby-adventures/grimms"
)

var dbhost = os.Getenv("DB_HOST")
var dbuser = os.Getenv("DB_USER")
var dbbase = os.Getenv("DB_BASE")
var dbpswd = os.Getenv("DB_PASS")
var dbport = os.Getenv("DB_PORT")

var CookieKey = []byte(os.Getenv("COO_KEY"))

var SupportLink = os.Getenv("SUPPORT_LINK")
var Botcolor = getEnvInt("BOT_COLOR")

var AppID = os.Getenv("APP_ID")
var DiscordSecret = os.Getenv("DISCORD_SECRET")
var Token = os.Getenv("DISCORD_TOKEN")
var ReportChannel = os.Getenv("REPORT_CHANNEL_ID")

// All characters
var BaseCharacters []chars.CharacterStruct
var BaseGrimms []grimms.GrimmStruct
