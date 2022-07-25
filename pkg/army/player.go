package army

import (
	"fmt"
	"sort"

	"github.com/spencerdrak/risk-calculator/pkg/util"
)

type Army struct {
	Owner    string
	ArmySize int
}

func (army Army) Attack() []int {
	roll := util.RollDice(true)

	sort.Sort(sort.Reverse(sort.IntSlice(roll)))

	switch army.ArmySize {
	case 0:
		roll[0] = 0
		roll[1] = 0
		roll[2] = 0
	case 1:
		roll[1] = 0
		roll[2] = 0
	case 2:
		roll[2] = 0
	}

	roll = roll[0:2]

	fmt.Printf("%s rolls: %v\n", army.Owner, roll)
	return roll
}

func (army Army) Defend() []int {
	roll := util.RollDice(true)

	sort.Sort(sort.Reverse(sort.IntSlice(roll)))

	switch army.ArmySize {
	case 0:
		roll[0] = 0
		roll[1] = 0
		roll[2] = 0
	case 1:
		roll[1] = 0
		roll[2] = 0
	default:
		roll[2] = 0
	}

	roll = roll[0:2]

	fmt.Printf("%s rolls: %v\n", army.Owner, roll)
	return roll
}
