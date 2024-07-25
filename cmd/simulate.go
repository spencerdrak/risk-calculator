/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spencerdrak/risk-calculator/pkg/army"
	"github.com/spencerdrak/risk-calculator/pkg/risk"
	"github.com/spf13/cobra"
)

// simulateCmd represents the simulate command
var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Test the program with large simulations",
	Long:  `Easy.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("USAGE: risk-calculator simulate <attacker count> <defender count>")
		}

		_, err1 := strconv.Atoi(args[0])
		_, err2 := strconv.Atoi(args[1])
		_, err3 := strconv.Atoi(args[2])
		if err1 == nil || err2 == nil || err3 == nil {
			return nil
		}
		return errors.New("arguments at position 0, 1 or 2 not convertable to int")
	},
	Run: func(cmd *cobra.Command, args []string) {
		iteration := 1
		wins := make(map[string]int)
		wins["attacker"] = 0
		wins["defender"] = 0
		simSize, _ := strconv.Atoi(args[0])
		attackerSize, _ := strconv.Atoi(args[1])
		defenderSize, _ := strconv.Atoi(args[2])
		for iteration < simSize {
			if iteration%100 == 0 {
				fmt.Printf("Iteration %d\n", iteration)
			}
			iteration += 1
			attacker := army.Army{
				ArmySize: attackerSize,
				Owner:    "attacker",
			}
			defender := army.Army{
				ArmySize: defenderSize,
				Owner:    "defender",
			}
			_, defender, err := risk.Run(attacker, defender)
			if err != nil {
				fmt.Printf("ERROR: %v\n, exiting...", err)
			}
			if defender.ArmySize == 0 {
				wins["attacker"] = wins["attacker"] + 1
			} else {
				wins["defender"] = wins["defender"] + 1
			}
		}
		fmt.Printf("Attacker wins: %f percent \n", ((float64(wins["attacker"]) / float64(iteration)) * 100))
		fmt.Printf("Defender wins: %f percent \n", ((float64(wins["defender"]) / float64(iteration)) * 100))
	},
}

func init() {
	rootCmd.AddCommand(simulateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// simulateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// simulateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
