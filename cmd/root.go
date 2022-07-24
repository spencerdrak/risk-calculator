/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/spencerdrak/risk-calculator/pkg/army"
	"github.com/spencerdrak/risk-calculator/pkg/util"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "risk-calculator",
	Short: "A calculator to play Risk, but faster",
	Long: `This calculator does a bunch of stuff, but mainly, you say Bob attacks Alice with 
	X number of pieces and Alice is defending with Y pieces. It'll assume you want to fight til one person is out of 
	pieces.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 4 {
			return errors.New("USAGE: risk-calculator attackername 12 defendername 24")
		}

		_, err1 := strconv.Atoi(args[1])
		_, err2 := strconv.Atoi(args[3])
		if err1 == nil || err2 == nil {
			return nil
		}
		return errors.New("arguments at position 1 or 3 not convertable to int")
	},
	Run: func(cmd *cobra.Command, args []string) {
		attackerSize, _ := strconv.Atoi(args[1])
		defenderSize, _ := strconv.Atoi(args[3])

		attacker := army.Army{
			ArmySize: attackerSize,
			Owner:    "Spencer",
		}
		defender := army.Army{
			ArmySize: defenderSize,
			Owner:    "John",
		}
		for attacker.ArmySize > 1 && defender.ArmySize > 0 {
			rand.Seed(time.Now().UnixNano())
			attack := attacker.Attack()
			defense := defender.Defend()
			comp, _ := util.CompareRolls(attack, defense)
			attacker.ArmySize = attacker.ArmySize - comp.AttackerLoss
			defender.ArmySize = defender.ArmySize - comp.DefenderLoss
			fmt.Println("-------------------------------------------")
		}

		fmt.Printf("Final Standings - Attacker: %d, Defender: %d", attacker.ArmySize, defender.ArmySize)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.risk-calculator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
