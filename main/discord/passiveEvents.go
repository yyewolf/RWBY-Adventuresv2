package discord

type PassiveFunc func(*CmdContext)

var RegisteredPassives []PassiveFunc
