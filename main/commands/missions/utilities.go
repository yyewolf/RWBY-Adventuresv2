package commands_missions

import namegenerator "github.com/yyewolf/Name-Generator"

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

func winMessages(i int) string {
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
