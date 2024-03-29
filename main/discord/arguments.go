package discord

import (
	"errors"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type CommandArg struct {
	Name  string
	Value interface{}
	Found bool
	Raw   *discordgo.ApplicationCommandInteractionDataOption
}

type Args []*CommandArg

type Arg struct {
	Name        string
	Description string
	Size        int
	Required    bool
	Choices     []*Choice
	Type        discordgo.ApplicationCommandOptionType
}

func (a *Args) GetArg(name string, index int, def interface{}) (argument *CommandArg) {
	s := *a
	argument = &CommandArg{
		Name:  name,
		Value: def,
		Found: false,
	}
	for i, arg := range s {
		if arg.Name == name {
			arg.Found = true
			return arg
		}
		if i == index {
			arg.Found = true
			argument = arg
		}
	}
	return
}

func (a *CommandArg) CharGrimmParse() (grimm bool, index int, err error) {
	if a == nil {
		return false, 0, errors.New("nil pointer")
	}
	switch a.Value.(type) {
	case string:
		break
	default:
		return false, 0, errors.New("not a string")
	}
	txt := a.Value.(string)
	txt = strings.ToLower(txt)
	if strings.Contains(txt, "g") {
		txt = strings.Replace(txt, "g", "", 1)
		grimm = true
	} else {
		txt = strings.Replace(txt, "c", "", 1)
	}
	index, err = strconv.Atoi(txt)
	return
}
