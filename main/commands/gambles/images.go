package commands_gambles

import (
	"bytes"
	"image"
	"image/png"
	"rwby-adventures/main/gambles"
	"rwby-adventures/main/static"
	"rwby-adventures/microservices"

	"github.com/fogleman/gg"
	uuid "github.com/satori/go.uuid"
)

var gambleImages map[string]image.Image

func init() {
	gambleImages = make(map[string]image.Image)
	BackgroundData, err := static.DatabaseFS.ReadFile("database/gamble/base.png")
	if err != nil {
		panic(err)
	}
	BackgroundDecoded, _, err := image.Decode(bytes.NewBuffer(BackgroundData))
	if err != nil {
		panic(err)
	}
	RLBData, err := static.DatabaseFS.ReadFile("database/gamble/rarelootbox.png")
	if err != nil {
		panic(err)
	}
	RLBDecoded, _, err := image.Decode(bytes.NewBuffer(RLBData))
	if err != nil {
		panic(err)
	}
	LBData, err := static.DatabaseFS.ReadFile("database/gamble/lootbox.png")
	if err != nil {
		panic(err)
	}
	LBDecoded, _, err := image.Decode(bytes.NewBuffer(LBData))
	if err != nil {
		panic(err)
	}
	ARMData, err := static.DatabaseFS.ReadFile("database/gamble/arm.png")
	if err != nil {
		panic(err)
	}
	ARMDecoded, _, err := image.Decode(bytes.NewBuffer(ARMData))
	if err != nil {
		panic(err)
	}
	CHARData, err := static.DatabaseFS.ReadFile("database/gamble/character.png")
	if err != nil {
		panic(err)
	}
	CHARDecoded, _, err := image.Decode(bytes.NewBuffer(CHARData))
	if err != nil {
		panic(err)
	}
	MONEYData, err := static.DatabaseFS.ReadFile("database/gamble/money.png")
	if err != nil {
		panic(err)
	}
	MONEYDecoded, _, err := image.Decode(bytes.NewBuffer(MONEYData))
	if err != nil {
		panic(err)
	}
	XPData, err := static.DatabaseFS.ReadFile("database/gamble/xp.png")
	if err != nil {
		panic(err)
	}
	XPDecoded, _, err := image.Decode(bytes.NewBuffer(XPData))
	if err != nil {
		panic(err)
	}
	LOSEData, err := static.DatabaseFS.ReadFile("database/gamble/lose.png")
	if err != nil {
		panic(err)
	}
	LOSEDecoded, _, err := image.Decode(bytes.NewBuffer(LOSEData))
	if err != nil {
		panic(err)
	}
	gambleImages["base"] = BackgroundDecoded
	gambleImages["rarelootbox"] = RLBDecoded
	gambleImages["lootbox"] = LBDecoded
	gambleImages["arm"] = ARMDecoded
	gambleImages["char"] = CHARDecoded
	gambleImages["money"] = MONEYDecoded
	gambleImages["xp"] = XPDecoded
	gambleImages["lose"] = LOSEDecoded
}

func createGambleImage(loots ...string) (r []string) {
	image := gg.NewContext(713, 403)

	//DRAW BASE
	image.DrawImage(gambleImages["base"], 0, 0)
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, image.Image())
	if err != nil {
		return nil
	}

	UUID := uuid.NewV4().String()
	_, err = gambles.UploadImage(&microservices.GambleUpload{
		UUID:  UUID,
		Image: buffer.Bytes(),
	})
	if err != nil {
		return nil
	}
	r = append(r, UUID)

	//FIRST LOOT
	image.DrawImage(gambleImages[loots[0]], 71, 84)
	buffer = new(bytes.Buffer)
	err = png.Encode(buffer, image.Image())
	if err != nil {
		return nil
	}

	UUID = uuid.NewV4().String()
	_, err = gambles.UploadImage(&microservices.GambleUpload{
		UUID:  UUID,
		Image: buffer.Bytes(),
	})
	if err != nil {
		return nil
	}
	r = append(r, UUID)

	//SECOND LOOT
	image.DrawImage(gambleImages[loots[1]], 268, 84)
	buffer = new(bytes.Buffer)
	err = png.Encode(buffer, image.Image())
	if err != nil {
		return nil
	}

	UUID = uuid.NewV4().String()
	_, err = gambles.UploadImage(&microservices.GambleUpload{
		UUID:  UUID,
		Image: buffer.Bytes(),
	})
	if err != nil {
		return nil
	}
	r = append(r, UUID)

	//THIRD LOOT
	image.DrawImage(gambleImages[loots[2]], 463, 84)
	buffer = new(bytes.Buffer)
	err = png.Encode(buffer, image.Image())
	if err != nil {
		return nil
	}

	UUID = uuid.NewV4().String()
	_, err = gambles.UploadImage(&microservices.GambleUpload{
		UUID:  UUID,
		Image: buffer.Bytes(),
	})
	if err != nil {
		return nil
	}
	r = append(r, UUID)

	return r
}
