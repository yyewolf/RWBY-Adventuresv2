package config

import (
	chars "rwby-adventures/characters"
	"rwby-adventures/grimms"
	"strings"
)

func LoadCharacters() {
	BaseCharacters = chars.BaseCharacters

	for i := range BaseCharacters {
		nameu := strings.ReplaceAll(BaseCharacters[i].Name, " ", "_")
		name := BaseCharacters[i].Name
		parent := BaseCharacters[i].Parent
		if parent != "" {
			nameu = strings.ReplaceAll(BaseCharacters[i].Parent, " ", "_")
			name = BaseCharacters[i].Parent
		}
		skin := strings.ReplaceAll(BaseCharacters[i].Skin, " ", "_")
		if skin == "" {
			skin = "Default"
		}
		BaseCharacters[i].IconURL = "https://img.rwbyadventures.com/" + nameu + "/" + skin + "_Icon.webp"
		BaseCharacters[i].WinURL = "https://img.rwbyadventures.com/" + nameu + "/" + skin + "_Win.gif"
		BaseCharacters[i].ImageFile = "database/images/" + name + "/" + skin + ".png"
	}

	BaseGrimms = grimms.BaseGrimms

	for i := range BaseGrimms {
		nameu := strings.ReplaceAll(BaseGrimms[i].Name, " ", "_")
		name := BaseGrimms[i].Name
		parent := BaseGrimms[i].Parent
		if parent != "" {
			nameu = strings.ReplaceAll(BaseGrimms[i].Parent, " ", "_")
			name = BaseGrimms[i].Parent
		}
		skin := strings.ReplaceAll(BaseGrimms[i].Skin, " ", "_")
		if skin == "" {
			skin = "Default"
		}
		BaseGrimms[i].IconURL = "https://img.rwbyadventures.com/" + nameu + "/" + skin + "_Icon.webp"
		BaseGrimms[i].WinURL = "https://img.rwbyadventures.com/" + nameu + "/" + skin + "_Win.gif"
		BaseGrimms[i].ImageFile = "database/grimms/" + name + "/" + skin + ".png"
	}
}
