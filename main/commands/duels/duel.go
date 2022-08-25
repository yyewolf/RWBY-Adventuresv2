package commands_duels

import (
	"fmt"
	"rwby-adventures/config"
	"rwby-adventures/main/discord"
	"rwby-adventures/models"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func duelAttack(ctx *discord.CmdContext) {
	defer func() {
		if err := recover(); err != nil {
			// Avoid crashes ?
			fmt.Println(err)
		}
	}()

	d := ctx.Menu.Data.(*BattleStruct)

	if _, ok := d.PlayersID[ctx.Author.ID]; !ok {
		return
	}
	AttackNumber := btnToAtk(strings.Split(ctx.ComponentData.CustomID, "-")[1])
	if AttackNumber == -1 {
		return
	}
	if AttackNumber == 100 {
		d.UseSemblance(ctx)
		return
	}

	pIndex := d.PlayersID[ctx.Author.ID]
	atk := d.Chars[pIndex].atks[AttackNumber]

	if d.TurnNumber+1-atk.LastUsed >= atk.Every {
		ctx.Reply(discord.ReplyParams{
			Content: &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredMessageUpdate,
			},
		})
		d.Players[pIndex].Chose = true
		d.Players[pIndex].Attack = AttackNumber
	} else {
		ctx.Reply(discord.ReplyParams{
			Content:   fmt.Sprintf("%s, you can use this again in %d turn(s).", ctx.Author.Username, atk.Every-(d.TurnNumber+1-atk.LastUsed)),
			Ephemeral: true,
		})
		return
	}

	if d.bothChose() {
		d.battleNextTurn(ctx)
	}
}

func (d *BattleStruct) UseSemblance(ctx *discord.CmdContext) {
	defer func() {
		if r := recover(); r != nil {
			// Avoid crashes ?
			fmt.Println(r)
		}
	}()

	pIndex := d.PlayersID[ctx.Author.ID]
	if d.Chars[pIndex].priority != 0 {
		if d.TurnNumber+1-d.Chars[pIndex].semblance.CustomData["lastTime"].(int) >= d.Chars[pIndex].semblance.Every {
			ctx.Reply(discord.ReplyParams{
				Content: &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseDeferredMessageUpdate,
				},
			})
			d.Players[pIndex].Chose = true
			d.Players[pIndex].Semblance = true
		} else {
			ctx.Reply(discord.ReplyParams{
				Content:   fmt.Sprintf("%s, you can use this again in %d turn(s).", ctx.Author.Username, d.Chars[pIndex].semblance.Every-(d.TurnNumber+1-d.Chars[pIndex].semblance.CustomData["lastTime"].(int))),
				Ephemeral: true,
			})
			return
		}
	}

	if d.bothChose() {
		d.battleNextTurn(ctx)
	}
}

func (d *BattleStruct) battleNextTurn(ctx *discord.CmdContext) {
	d.TurnNumber++
	p1 := d.Players[0]
	p2 := d.Players[1]
	c1 := d.Chars[0]
	c2 := d.Chars[1]
	p1atk := d.Chars[0].atks[d.Players[0].Attack]
	p2atk := d.Chars[1].atks[d.Players[1].Attack]

	//calls passive semblance
	c1.semblance.callPassive(c1.stats)
	c2.semblance.callPassive(c2.stats)

	LatestEvent := discordgo.MessageEmbedField{
		Name: "Latest event :",
	}

	var Order []int

	if p1atk.Speed > p2atk.Speed {
		Order = []int{0, 1}
	} else {
		Order = []int{1, 0}
	}

	if p1.Semblance && p2.Semblance && c1.priority > c2.priority {
		Order = []int{0, 1}
	} else if p1.Semblance && p2.Semblance && c1.priority < c2.priority {
		Order = []int{1, 0}
	} else if p1.Semblance {
		Order = []int{0, 0}
	} else if p2.Semblance {
		Order = []int{1, 0}
	}

	for CurrentPlayer := range Order {
		p := d.Players[CurrentPlayer]
		op := d.Players[opponentPlayer(CurrentPlayer)]
		oc := d.Chars[opponentPlayer(CurrentPlayer)]

		damageDealt, killed, stunned, healed, dodged := 0, false, false, 0, false

		//Do normal attack if no semblance
		if !p.Semblance {
			damageDealt, killed, stunned, healed, dodged = oc.TakeDamage(d.Chars[CurrentPlayer], d.Chars[CurrentPlayer].atks[p.Attack])
			if !d.Chars[CurrentPlayer].stunned {
				d.Chars[CurrentPlayer].atks[p.Attack].LastUsed = d.TurnNumber
			}
		} else {
			d.Players[CurrentPlayer].Semblance = false
			f := d.Chars[CurrentPlayer].semblance.callMain(d.Chars[CurrentPlayer].stats)
			damageDealt, killed, stunned, healed, dodged = oc.TakeDamage(d.Chars[CurrentPlayer], &BattleAtks{
				StunChance: f.StunChance,
				Damages:    f.Damage,
				Heal:       f.Heal,
			})
			if !d.Chars[CurrentPlayer].stunned {
				d.Chars[CurrentPlayer].semblance.CustomData["lastTime"] = d.TurnNumber
			}
		}

		//Makes semblances calls
		d.Chars[CurrentPlayer].semblance.callAttacked(d.Chars[CurrentPlayer].stats, damageDealt)
		d.Chars[opponentPlayer(CurrentPlayer)].semblance.callGotAttacked(d.Chars[opponentPlayer(CurrentPlayer)].stats, damageDealt)

		if damageDealt > 0 {
			LatestEvent.Value += p.User.Username + " dealt " + strconv.Itoa(damageDealt) + "HP to " + op.User.Username + "\n"
		}
		if healed > 0 {
			LatestEvent.Value += p.User.Username + " healed " + strconv.Itoa(healed) + "HP !\n"
		}
		if stunned {
			LatestEvent.Value += p.User.Username + " has stunned " + op.User.Username + "!\n"
		}
		if dodged {
			LatestEvent.Value += op.User.Username + " has dodged " + p.User.Username + "!\n"
		}

		if killed && !d.Finished {
			LatestEvent.Value += p.User.Username + " has defeated " + op.User.Username + "!\n"
			d.Finished = true
			d.Won = p
			d.Lost = op
			d.WonImg = d.Chars[CurrentPlayer].winURL
			break
		}
	}

	if !d.Finished {
		LatestEvent.Value += "Prepare your attacks for the next round!"
	} else {
		LatestEvent.Value += "This was the last turn, congrats !"
	}

	d.Original.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       p1.User.Username + " VS " + p2.User.Username,
			Description: "Turn : " + strconv.Itoa(d.TurnNumber),
			Fields: []*discordgo.MessageEmbedField{
				&LatestEvent,
				{
					Name:   p1.User.Username + "'s HP :",
					Value:  strconv.Itoa(c1.stats.Health),
					Inline: true,
				},
				{
					Name:   p2.User.Username + "'s HP :",
					Value:  strconv.Itoa(c2.stats.Health),
					Inline: true,
				},
			},
			Image: &discordgo.MessageEmbedImage{
				URL: "attachment://vs.png",
			},
			Color: config.Botcolor,
		},
		Edit: true,
	})

	if d.Finished {
		d.End(ctx)
		return
	}

	p1.Chose = false
	p2.Chose = false
}

func (d *BattleStruct) End(ctx *discord.CmdContext) {
	WinXP := duelXP(ctx, d.Won, true)
	LostXP := duelXP(ctx, d.Lost, false)
	ctx.Reply(discord.ReplyParams{
		Content: &discordgo.MessageEmbed{
			Title:       d.Won.User.Username + " won !",
			Description: fmt.Sprintf("Winner earned : %sXP\nOther player earned : %sXP", WinXP, LostXP),
			Image: &discordgo.MessageEmbedImage{
				URL: d.WonImg,
			},
			Color: config.Botcolor,
		},
	})
}

func duelXP(ctx *discord.CmdContext, player *BattlePlayer, winner bool) (XP string) {
	p := models.GetPlayer(player.User.ID)
	if p.SelectedID != player.SelectedID {
		p.SelectedID = player.SelectedID
		p.SelectedType = player.SelectedType
		if p.SelectedType == models.GrimmType {
			for _, g := range p.Grimms {
				if g.GrimmID == p.SelectedID {
					p.SelectedGrimm = g
					break
				}
			}
		} else {
			for _, c := range p.Characters {
				if c.CharID == p.SelectedID {
					p.SelectedChar = c
					break
				}
			}
		}
	}
	if !p.SelectedChar.Valid() && !p.SelectedGrimm.Valid() {
		return
	}

	WillAdd := p.CalcSelectedXP(4, false)

	if !winner {
		WillAdd = WillAdd / 2
		p.Stats.BattlesLost++
	} else {
		p.Stats.BattlesWon++
	}
	p.Stats.Save()

	p.GiveSelectedXP(WillAdd)
	return strconv.FormatInt(WillAdd, 10)
}
