package discord

import (
	chars "rwby-adventures/characters"
	"rwby-adventures/config"
	"strings"
)

func LoadAllCharacters() {
	config.BaseCharacters = chars.BaseCharacters

	for i := range config.BaseCharacters {
		nameu := strings.ReplaceAll(config.BaseCharacters[i].Name, " ", "_")
		name := config.BaseCharacters[i].Name
		parent := config.BaseCharacters[i].Parent
		if parent != "" {
			nameu = strings.ReplaceAll(config.BaseCharacters[i].Parent, " ", "_")
			name = config.BaseCharacters[i].Parent
		}
		skin := strings.ReplaceAll(config.BaseCharacters[i].Skin, " ", "_")
		if skin == "" {
			skin = "Default"
		}
		config.BaseCharacters[i].IconURL = "https://img.rwbyadventures.com/" + nameu + "/" + skin + "_Icon.webp"
		config.BaseCharacters[i].WinURL = "https://img.rwbyadventures.com/" + nameu + "/" + skin + "_Win.gif"
		config.BaseCharacters[i].ImageFile = name + "/" + skin + ".png"
	}
}
