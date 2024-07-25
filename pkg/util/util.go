package util

import (
	"fmt"
	"math/rand"
	"time"
)

type RollComparison struct {
	AttackerLoss int
	DefenderLoss int
}

var print_due bool = false

func RollDice(attacker bool) []int {
	rand.Seed(time.Now().UnixNano())
	return []int{rand.Intn(6) + 1, rand.Intn(6) + 1, rand.Intn(6) + 1}
}

func CompareRolls(attackerRoll []int, defenderRoll []int) (RollComparison, error) {
	output := RollComparison{
		AttackerLoss: 0,
		DefenderLoss: 0,
	}

	for i := range attackerRoll {
		attack := attackerRoll[i]
		defense := defenderRoll[i]

		if attack == 0 || defense == 0 {
			continue
		}

		if attack > defense {
			output.DefenderLoss += 1
		} else {
			output.AttackerLoss += 1
		}
	}

	if print_due {
		fmt.Printf("Attacker loses: %d, Defender loses: %d\n", output.AttackerLoss, output.DefenderLoss)
	}

	return output, nil
}
