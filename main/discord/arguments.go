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

func (a *Args) GetArg(name string, index int) (*CommandArg, error) {
	s := *a
	var temp *CommandArg
	var fallback = true
	for i, arg := range s {
		if arg.Name != "" {
			fallback = false
		}
		if arg.Name == name {
			return arg, nil
		}
		if i == index {
			temp = arg
		}
	}
	if temp == nil || !fallback {
		return nil, errors.New("nil pointer")
	}
	return temp, nil
}

func (a *CommandArg) CharGrimmParse() (grimm bool, index int, err error) {
	if a == nil {
		return false, 0, errors.New("nil pointer")
	}
	txt := a.Value.(string)
	txt = strings.ToLower(txt)
	if strings.Contains(txt, "g") {
		txt = strings.Replace(txt, "g", "", 1)
		grimm = true
	}
	index, err = strconv.Atoi(txt)
	return
}
