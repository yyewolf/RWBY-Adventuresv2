package commands_missions

import (
	"math/rand"
	"rwby-adventures/config"

	namegenerator "github.com/yyewolf/Name-Generator"
)

const (
	missionSearchRescue = iota
	missionSearchDestroy
	missionPerimeterDef
	missionVillageSecu
	missionBounty
	missionEscort
	missionSummer
	missionTai
	missionRaven
	missionQrow
)

const (
	huntFindBaby = iota
	huntDestroyCity
	huntKillHero
	huntSalemTask
	huntTakeRest
	huntDestroyBeacon
)

var generator namegenerator.Session

func missionToString(i int) string {
	switch i {
	case missionSearchRescue:
		return "Search and Rescue"
	case missionSearchDestroy:
		return "Search and Destroy"
	case missionPerimeterDef:
		return "Perimeter Defense"
	case missionVillageSecu:
		return "Village Security"
	case missionBounty:
		return "Bounty"
	case missionEscort:
		return "Escort"
		//STRQ EVENT
		/*
			case missionSummer:
				return "Summer needs your help accomplishing her last mission"
			case missionTai:
				return "Tai needs your help cleaning stuff around the house"
			case missionRaven:
				return "Raven needs your help to move her camp"
			case missionQrow:
				return "Qrow has been missing for a few days, you should check on him"
		*/
	}
	return ""
}

func huntToString(i int) string {
	switch i {
	case huntFindBaby:
		return "Find a Baby Grimm"
	case huntDestroyCity:
		return "Destroy a Town"
	case huntKillHero:
		return "Kill Hero"
	case huntSalemTask:
		return "Salem Task"
	case huntTakeRest:
		return "Restore and Sleep"
	case huntDestroyBeacon:
		return "Destroy Beacon"
	}
	return ""
}

func missionWinMessages(i int) string {
	switch i {
	case missionSearchRescue:
		name, _ := generator.GetName()
		return "You successfully rescued **" + name + "** !\nYou earned : {mission.Earned}"
	case missionSearchDestroy:
		return "You successfully cleared out the Grimms !\nYou earned : {mission.Earned}"
	case missionPerimeterDef:
		return "You successfully defended the kingdom !\nYou earned : {mission.Earned}"
	case missionVillageSecu:
		name, _ := generator.GetCityName()
		return "You successfully defended **" + name + "** !\nYou earned : {mission.Earned}"
	case missionBounty:
		name, _ := generator.GetName()
		return "You successfully captured **" + name + "** !\nYou earned : {mission.Earned}"
	case missionEscort:
		name, _ := generator.GetName()
		return "You successfully escorted **" + name + "** !\nYou earned : {mission.Earned}"
		//STRQ EVENT
		/*
			case missionSummer:
				return "Summer accepted to join you on your journey."
			case missionTai:
				return "Tai accepted to join you on your journey."
			case missionRaven:
				return "Raven accepted to join you on your journey."
			case missionQrow:
				return "Qrow accepted to join you on your journey."
		*/
	}
	return ""
}

func huntWinMessages(i int) string {
	switch i {
	case huntFindBaby:
		name, _ := generator.GetName()
		return "You successfully found **" + name + "** Grimm !\nYou earned : {hunt.Earned}"
	case huntDestroyCity:
		name, _ := generator.GetCityName()
		return "You successfully destroyed **" + name + "** !\nYou earned : {hunt.Earned}"
	case huntKillHero:
	pick:
		i := rand.Intn(len(config.BaseCharacters))
		if config.BaseCharacters[i].Limited {
			goto pick
		}
		name := config.BaseCharacters[i].Name
		return "You successfully killed **" + name + "** !\nYou earned : {hunt.Earned}"
	case huntSalemTask:
		return "You successfully achieved what Salem told you to do !\nYou earned : {hunt.Earned}"
	case huntTakeRest:
		return "You successfully **took a rest** !\nYou earned : {hunt.Earned}"
	case huntDestroyBeacon:
		return "You successfully destroyed **Beacon** !\nYou earned : {hunt.Earned}"
	}
	return ""
}
