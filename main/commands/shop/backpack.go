package commands_shop

import (
	"errors"
	"rwby-adventures/models"
)

func buyBackpack(p *models.Player) (string, error) {
	BPrice := int64(12000*((p.Shop.Extensions)+1) + 5000)
	if p.TotalBalance() < BPrice {
		return "You don't have enough money.", errors.New("err")
	}
	p.Shop.Extensions++
	p.Balance -= BPrice

	p.Save()
	p.Shop.Save()
	return "You bought a backpack extension !", nil
}
