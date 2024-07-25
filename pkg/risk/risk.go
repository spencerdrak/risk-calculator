package risk

import (
	"fmt"

	"github.com/spencerdrak/risk-calculator/pkg/army"
	"github.com/spencerdrak/risk-calculator/pkg/util"
)

var print_due bool = false

func Run(attacker army.Army, defender army.Army) (army.Army, army.Army, error) {
	for attacker.ArmySize > 1 && defender.ArmySize > 0 {
		attack := attacker.Attack()
		defense := defender.Defend()
		comp, err := util.CompareRolls(attack, defense)
		if err != nil {
			return attacker, defender, err
		}
		attacker.ArmySize = attacker.ArmySize - comp.AttackerLoss
		defender.ArmySize = defender.ArmySize - comp.DefenderLoss
		if print_due {
			fmt.Println("-------------------------------------------")
		}
	}
	return attacker, defender, nil
}
