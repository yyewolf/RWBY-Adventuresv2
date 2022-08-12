package websocket

import (
	"math/rand"
	chars "rwby-adventures/characters"
	"rwby-adventures/config"
	"rwby-adventures/grimms"

	"github.com/ambelovsky/gosf"
)

func getRandomPersonas(client *gosf.Client, request *gosf.Request) *gosf.Message {
	// Get 3 random grimms & 3 random chars :
	var ids []int
	var g []*grimms.GrimmStruct
	for i := 0; i < 3; i++ {
	reroll1:
		rng := rand.Intn(len(config.BaseGrimms))
		// Make sure we don't get the same one twice :
		for _, id := range ids {
			if id == rng {
				goto reroll1
			}
		}

		ids = append(ids, rng)
		g = append(g, &config.BaseGrimms[rng])
	}

	ids = []int{}
	var c []*chars.CharacterStruct
	for i := 0; i < 3; i++ {
	reroll2:
		rng := rand.Intn(len(config.BaseCharacters))
		// Make sure we don't get the same one twice :
		for _, id := range ids {
			if id == rng {
				goto reroll2
			}
		}

		ids = append(ids, rng)
		c = append(c, &config.BaseCharacters[rng])
	}
	msg := gosf.NewSuccessMessage()
	msg.Body = make(map[string]interface{})
	msg.Body["grimms"] = g
	msg.Body["characters"] = c
	return msg
}
