package commands_shop

import (
	"errors"
	"rwby-adventures/models"
)

func buyLootbox(p *models.Player) (string, error) {
	if p.TotalBalance() < 1750 {
		return "You don't have enough money.", errors.New("err")
	}
	if p.Lootbox()+2 > p.MaxChar() {
		return "You already have too much loot boxes.", errors.New("err")
	}
	p.Balance -= 1750
	p.Boxes.Boxes += 2

	p.Save()
	p.Boxes.Save()
	return "You bought two loot boxes.", nil
}

func buyGrimmbox(p *models.Player) (string, error) {
	if p.TotalBalance() < 2500 {
		return "You don't have enough money.", errors.New("err")
	}
	if p.Lootbox()+2 > p.MaxChar() {
		return "You already have too much loot boxes.", errors.New("err")
	}
	p.Balance -= 2500
	p.Boxes.GrimmBoxes += 2

	p.Save()
	p.Boxes.Save()
	return "You bought two grimm boxes.", nil
}

func buyRarelootbox(p *models.Player) (string, error) {
	if p.TotalBalance() < 5000 {
		return "You don't have enough money.", errors.New("err")
	}
	if p.Lootbox()+3 > p.MaxChar() {
		return "You already have too much loot boxes.", errors.New("err")
	}
	p.Balance -= 5000
	p.Boxes.RareBoxes += 3

	p.Save()
	p.Boxes.Save()
	return "You bought 3 rare loot boxes.", nil
}

func buyRareGrimmbox(p *models.Player) (string, error) {
	if p.TotalBalance() < 7500 {
		return "You don't have enough money.", errors.New("err")
	}
	if p.Lootbox()+3 > p.MaxChar() {
		return "You already have too much loot boxes.", errors.New("err")
	}
	p.Balance -= 5000
	p.Boxes.RareGrimmBoxes += 3

	p.Save()
	p.Boxes.Save()
	return "You bought 3 rare grimm boxes.", nil
}

func exchangeRareGrimmbox(p *models.Player) (string, error) {
	if p.Boxes.GrimmBoxes < 10 {
		return "You don't have enough grimm boxes.", errors.New("err")
	}
	p.Boxes.GrimmBoxes -= 10
	p.Boxes.RareGrimmBoxes += 1

	p.Boxes.Save()
	return "You bought 1 rare grimm box.", nil
}

func exchangeRareLootbox(p *models.Player) (string, error) {
	if p.Boxes.Boxes < 10 {
		return "You don't have enough loot boxes.", errors.New("err")
	}
	p.Boxes.Boxes -= 10
	p.Boxes.RareBoxes += 1

	p.Boxes.Save()
	return "You bought 1 rare loot box.", nil
}
