package config

import (
	chars "rwby-adventures/characters"
	"rwby-adventures/grimms"
)

//var dbhost = "admin.rwbyadventures.com"
//var dbuser = "yewolf"
//var dbpswd = "ftT6A4MrF6hPt"

var dbhost = "127.0.0.1"
var dbport = "5432"
var dbuser = "postgres"
var dbpswd = "admin"

var SupportLink = "https://discord.gg/adJGyVxv7H"
var Botcolor = 3859607

var AppID = "375700234120200194"
var Token = "Mzc1NzAwMjM0MTIwMjAwMTk0.WftYGw.7Dpvcfk1sBR2QksWuV-x4IbpziI"

// All characters
var BaseCharacters []chars.CharacterStruct
var BaseGrimms []grimms.GrimmStruct
