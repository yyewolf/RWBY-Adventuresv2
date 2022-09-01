package websocket

import (
	"encoding/json"
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/yyewolf/gosf"
)

type WebUser struct {
	Name string
	ID   string
}

type TradeTemplateData struct {
	User    WebUser
	Token   string
	OtherID string
	Host    string
}

type TradeClientRequest struct {
	Action string `json:"action"`
	Target string `json:"target"`
}

type Persona struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Rarity   string  `json:"rarity"`
	ImageURL string  `json:"imageurl"`
}

type Misc struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Max      int64  `json:"max"`
	ImageURL string `json:"imageurl"`
}

func TradeConnect(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		return gosf.NewFailureMessage("f")
	}
	d := data.(*TradeTemplateData)
	client.Join(d.Token)
	return gosf.NewSuccessMessage()
}

func TradeInfos(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		return gosf.NewFailureMessage("f")
	}
	d := data.(*TradeTemplateData)
	raw, ok := request.Message.Body["data"]
	if !ok {
		return gosf.NewFailureMessage("f")
	}
	body, err := json.Marshal(raw)
	if err != nil {
		return gosf.NewFailureMessage("f")
	}
	content := &TradeClientRequest{}
	err = json.Unmarshal(body, content)
	if err != nil {
		return gosf.NewFailureMessage("f")
	}

	if content.Target != d.OtherID && content.Target != d.User.ID {
		return gosf.NewFailureMessage("f")
	}

	p := models.GetPlayer(content.Target)

	msg := gosf.NewSuccessMessage("")
	msg.Body = map[string]interface{}{}
	msg.Body["grimms"] = []Persona{}
	msg.Body["characters"] = []Persona{}
	msg.Body["misc"] = []Misc{}
	switch content.Action {
	case "grimms":
		var list []Persona
		fmt.Println("grimms: ", p.Grimms)
		for _, grimm := range p.Grimms {
			list = append(list, Persona{
				ID:       grimm.GrimmID,
				Name:     grimm.Name,
				Value:    grimm.Value,
				Rarity:   grimm.RarityString(),
				ImageURL: grimm.ToRealGrimm().IconURL,
				Type:     "grimm",
			})
		}
		msg.Body["grimms"] = list
		break
	case "chars":
		var list []Persona
		for _, char := range p.Characters {
			list = append(list, Persona{
				ID:       char.CharID,
				Name:     char.Name,
				Value:    char.Value,
				Rarity:   char.RarityString(),
				ImageURL: char.ToRealChar().IconURL,
				Type:     "char",
			})
		}
		msg.Body["characters"] = list
		break
	case "misc":
		var list []Misc
		list = append(list, Misc{
			ID:       "money",
			Name:     "Money",
			Max:      p.TotalBalance(),
			ImageURL: "https://cdn.shopify.com/s/files/1/1061/1924/products/Money_Bag_Emoji_large.png?v=1571606064",
		})
		list = append(list, Misc{
			ID:       "classicboxes",
			Name:     "Classic",
			Max:      int64(p.Boxes.Boxes),
			ImageURL: "https://freeiconshop.com/wp-content/uploads/edd/box-outline-filled.png",
		})
		list = append(list, Misc{
			ID:       "rareboxes",
			Name:     "Rare",
			Max:      int64(p.Boxes.RareBoxes),
			ImageURL: "https://cdn-icons-png.flaticon.com/512/1907/1907938.png",
		})
		list = append(list, Misc{
			ID:       "classicgrimmboxes",
			Name:     "Grimm",
			Max:      int64(p.Boxes.GrimmBoxes),
			ImageURL: "https://freeiconshop.com/wp-content/uploads/edd/box-outline-filled.png",
		})
		list = append(list, Misc{
			ID:       "raregrimmboxes",
			Name:     "Rare Grimm",
			Max:      int64(p.Boxes.RareGrimmBoxes),
			ImageURL: "https://cdn-icons-png.flaticon.com/512/1907/1907938.png",
		})
		msg.Body["misc"] = list
		break
	}
	msg.Body["target"] = p.DiscordID
	return msg
}

type TradeValidateContent struct {
	Characters     []string `json:"chars"`
	Grimms         []string `json:"grimms"`
	Money          int64    `json:"money"`
	Boxes          int64    `json:"classicboxes"`
	RareBoxes      int64    `json:"rareboxes"`
	GrimmBoxes     int64    `json:"classicgrimmboxes"`
	RareGrimmBoxes int64    `json:"raregrimmboxes"`
}

type TradeClientValidate struct {
	UserSends   TradeValidateContent `json:"own"`
	TargetSends TradeValidateContent `json:"other"`
}

func TradeValidate(client *gosf.Client, request *gosf.Request) *gosf.Message {
	data, found := GetToken(request)
	if !found {
		return gosf.NewFailureMessage("f")
	}
	d := data.(*TradeTemplateData)
	raw, ok := request.Message.Body["data"]
	if !ok {
		return gosf.NewFailureMessage("f")
	}
	body, err := json.Marshal(raw)
	if err != nil {
		return gosf.NewFailureMessage("f")
	}
	content := &TradeClientValidate{}
	err = json.Unmarshal(body, content)
	if err != nil {
		return gosf.NewFailureMessage("f")
	}

	player := models.GetPlayer(d.User.ID)
	target := models.GetPlayer(d.OtherID)

	if content.UserSends.Money < 0 || content.UserSends.Boxes < 0 || content.UserSends.RareBoxes < 0 || content.UserSends.GrimmBoxes < 0 || content.UserSends.RareGrimmBoxes < 0 {
		return gosf.NewFailureMessage("There is an issue with your boxes.")
	}
	if content.TargetSends.Money < 0 || content.TargetSends.Boxes < 0 || content.TargetSends.RareBoxes < 0 || content.TargetSends.GrimmBoxes < 0 || content.TargetSends.RareGrimmBoxes < 0 {
		return gosf.NewFailureMessage("There is an issue with their boxes.")
	}

	// We check that the player has what he claims to
	if !player.VerifyChars(content.UserSends.Characters) || !player.VerifyGrimms(content.UserSends.Grimms) {
		return gosf.NewFailureMessage("There is an issue with your personas.")
	}

	// We check that the target has what the player claims he has
	if !target.VerifyChars(content.TargetSends.Characters) || !target.VerifyGrimms(content.TargetSends.Grimms) {
		return gosf.NewFailureMessage("There is an issue with your target's personas.")
	}

	if player.TradeSent >= 5 {
		return gosf.NewFailureMessage("You have too much trade on hold.")
	}
	if target.TradeReceived >= 8 {
		return gosf.NewFailureMessage("This person already has too much trades on hold.")
	}

	PlayerResult := &TradeValidateContent{
		Money:          int64(player.TotalBalance()) + content.TargetSends.Money - content.UserSends.Money,
		Boxes:          int64(player.Boxes.Boxes) + content.TargetSends.Boxes - content.UserSends.Boxes,
		RareBoxes:      int64(player.Boxes.RareBoxes) + content.TargetSends.RareBoxes - content.UserSends.RareBoxes,
		GrimmBoxes:     int64(player.Boxes.GrimmBoxes) + content.TargetSends.GrimmBoxes - content.UserSends.GrimmBoxes,
		RareGrimmBoxes: int64(player.Boxes.RareGrimmBoxes) + content.TargetSends.RareGrimmBoxes - content.UserSends.RareGrimmBoxes,
	}

	TargetResult := &TradeValidateContent{
		Money:          int64(target.TotalBalance()) - content.TargetSends.Money + content.UserSends.Money,
		Boxes:          int64(target.Boxes.Boxes) - content.TargetSends.Boxes + content.UserSends.Boxes,
		RareBoxes:      int64(target.Boxes.RareBoxes) - content.TargetSends.RareBoxes + content.UserSends.RareBoxes,
		GrimmBoxes:     int64(target.Boxes.GrimmBoxes) - content.TargetSends.GrimmBoxes + content.UserSends.GrimmBoxes,
		RareGrimmBoxes: int64(target.Boxes.RareGrimmBoxes) - content.TargetSends.RareGrimmBoxes + content.UserSends.RareGrimmBoxes,
	}

	if PlayerResult.Money < 0 {
		return gosf.NewFailureMessage("You don't have enough balance.")
	}
	if PlayerResult.Boxes < 0 {
		return gosf.NewFailureMessage("You don't have enough boxes.")
	}
	if PlayerResult.RareBoxes < 0 {
		return gosf.NewFailureMessage("You don't have enough rare boxes.")
	}
	if PlayerResult.GrimmBoxes < 0 {
		return gosf.NewFailureMessage("You don't have enough grimm boxes.")
	}
	if PlayerResult.RareGrimmBoxes < 0 {
		return gosf.NewFailureMessage("You don't have enough rare grimm boxes.")
	}
	if PlayerResult.Boxes+PlayerResult.RareBoxes > int64(player.MaxChar()) {
		return gosf.NewFailureMessage("You would receive too much boxes.")
	}
	if PlayerResult.RareGrimmBoxes+PlayerResult.GrimmBoxes > int64(player.MaxChar()) {
		return gosf.NewFailureMessage("You would receive too much grimm boxes.")
	}

	if TargetResult.Money < 0 {
		return gosf.NewFailureMessage("They don't have enough balance.")
	}
	if TargetResult.Boxes < 0 {
		return gosf.NewFailureMessage("They don't have enough boxes.")
	}
	if TargetResult.RareBoxes < 0 {
		return gosf.NewFailureMessage("They don't have enough rare boxes.")
	}
	if TargetResult.GrimmBoxes < 0 {
		return gosf.NewFailureMessage("They don't have enough grimm boxes.")
	}
	if TargetResult.RareGrimmBoxes < 0 {
		return gosf.NewFailureMessage("They don't have enough rare grimm boxes.")
	}
	if TargetResult.Boxes+TargetResult.RareBoxes > int64(player.MaxChar()) {
		return gosf.NewFailureMessage("They would receive too much boxes.")
	}
	if TargetResult.RareGrimmBoxes+TargetResult.GrimmBoxes > int64(player.MaxChar()) {
		return gosf.NewFailureMessage("They would receive too much grimm boxes.")
	}

	if len(player.Characters)+len(content.TargetSends.Characters)-len(content.UserSends.Characters) > player.CharLimit {
		return gosf.NewFailureMessage("You would receive too much characters.")
	}
	if len(player.Grimms)+len(content.TargetSends.Grimms)-len(content.UserSends.Grimms) > player.CharLimit {
		return gosf.NewFailureMessage("You would receive too much grimms.")
	}
	if len(player.Characters)+len(content.TargetSends.Characters)-len(content.UserSends.Characters) < 0 {
		return gosf.NewFailureMessage("You would be left without any characters.")
	}
	if len(player.Grimms)+len(content.TargetSends.Grimms)-len(content.UserSends.Grimms) < 0 {
		return gosf.NewFailureMessage("You would be left without any grimms.")
	}

	if len(target.Characters)-len(content.TargetSends.Characters)+len(content.UserSends.Characters) > target.CharLimit {
		return gosf.NewFailureMessage("They would receive too much characters.")
	}
	if len(target.Grimms)-len(content.TargetSends.Grimms)+len(content.UserSends.Grimms) > target.CharLimit {
		return gosf.NewFailureMessage("They would receive too much grimms.")
	}
	if len(target.Characters)-len(content.TargetSends.Characters)+len(content.UserSends.Characters) < 0 {
		return gosf.NewFailureMessage("They would be left without any characters.")
	}
	if len(target.Grimms)-len(content.TargetSends.Grimms)+len(content.UserSends.Grimms) < 0 {
		return gosf.NewFailureMessage("They would be left without any grimms.")
	}

	s := discord.Session
	SenderUser, _ := s.User(player.DiscordID)
	SenderDM, _ := s.UserChannelCreate(player.DiscordID)
	ReceiverUser, _ := s.User(target.DiscordID)
	ReceiverDM, _ := s.UserChannelCreate(target.DiscordID)

	SenderMessage := content.UserSends.PrepareMessage(player)
	ReceiverMessage := content.TargetSends.PrepareMessage(target)

	TradeID := fmt.Sprintf("%d", time.Now().UnixMilli())
	t := models.Trade{
		ID:          TradeID,
		SenderID:    player.DiscordID,
		ReceiverID:  target.DiscordID,
		StartedAt:   time.Now().Unix(),
		UserSends:   content.UserSends.ToModel(TradeID, player.DiscordID, "sender"),
		TargetSends: content.TargetSends.ToModel(TradeID, target.DiscordID, "receiver"),
	}

	//Level 5 Common Ruby
	FieldLeft := &discordgo.MessageEmbedField{
		Name:   SenderUser.String() + " is offering :",
		Value:  "`" + SenderMessage + "`",
		Inline: true,
	}
	FieldRight := &discordgo.MessageEmbedField{
		Name:   ReceiverUser.String() + " is offering :",
		Value:  "`" + ReceiverMessage + "`",
		Inline: true,
	}

	SenderEmbed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Trade ID `%s` recap :", TradeID),
		Description: fmt.Sprintf("Type : `%strade delete %s` to cancel it.", discord.CommandRouter.Prefix, TradeID),
		Fields:      []*discordgo.MessageEmbedField{FieldLeft, FieldRight},
		Color:       config.Botcolor,
	}

	ReceiverEmbed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Trade ID `%s` recap :", TradeID),
		Description: fmt.Sprintf("Type : `%strade refuse %s` to refuse it.\nType : `%strade accept %s` to accept it.", discord.CommandRouter.Prefix, TradeID, discord.CommandRouter.Prefix, TradeID),
		Fields:      []*discordgo.MessageEmbedField{FieldLeft, FieldRight},
		Color:       config.Botcolor,
	}
	s.ChannelMessageSendEmbed(SenderDM.ID, SenderEmbed)
	s.ChannelMessageSendEmbed(ReceiverDM.ID, ReceiverEmbed)
	config.Database.Create(t)
	config.Database.Create(t.TargetSends)
	config.Database.Create(t.UserSends)
	return gosf.NewSuccessMessage("gg")
}

func (c *TradeValidateContent) PrepareMessage(p *models.Player) string {
	Message := ""
	if c.Money > 0 {
		Message += fmt.Sprintf("Money : %dⱠ.\n", c.Money)
	}
	if c.Boxes > 0 {
		Message += fmt.Sprintf("Boxes : %dⱠ.\n", c.Boxes)
	}
	if c.GrimmBoxes > 0 {
		Message += fmt.Sprintf("Grimm Boxes : %dⱠ.\n", c.GrimmBoxes)
	}
	if c.RareBoxes > 0 {
		Message += fmt.Sprintf("Rare Boxes : %dⱠ.\n", c.RareBoxes)
	}
	if c.RareGrimmBoxes > 0 {
		Message += fmt.Sprintf("Rare Grimm Boxes : %dⱠ.\n", c.RareGrimmBoxes)
	}
	for i, CharID := range c.Characters {
		current := &models.Character{}
		for _, char := range p.Characters {
			if char.CharID == CharID {
				current = char
				break
			}
		}
		if current.Name == "" {
			continue
		}
		if i != 0 {
			Message += "\n"
		}
		Message += current.MiniString()
	}
	for i, GrimmID := range c.Grimms {
		current := &models.Grimm{}
		for _, grimm := range p.Grimms {
			if grimm.GrimmID == GrimmID {
				current = grimm
				break
			}
		}
		if current.Name == "" {
			continue
		}
		if i != 0 {
			Message += "\n"
		}
		Message += current.MiniString()
	}
	return Message
}

func (c *TradeValidateContent) ToModel(TradeID string, UserID string, Type string) *models.TradeContent {
	return &models.TradeContent{
		TradeID:        TradeID,
		Type:           Type,
		UserID:         UserID,
		Characters:     c.Characters,
		Grimms:         c.Grimms,
		Money:          c.Money,
		Boxes:          c.Boxes,
		RareBoxes:      c.RareBoxes,
		GrimmBoxes:     c.GrimmBoxes,
		RareGrimmBoxes: c.RareGrimmBoxes,
	}
}
