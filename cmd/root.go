/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spencerdrak/risk-calculator/pkg/army"
	"github.com/spencerdrak/risk-calculator/pkg/risk"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "risk-calculator",
	Short: "A calculator to play Risk, but faster",
	Long: `This calculator does a bunch of stuff, but mainly, you say Bob attacks Alice with 
	X number of pieces and Alice is defending with Y pieces. It'll assume you want to fight til one person is out of 
	pieces. The attacker will go no lower than 1 piece, and the defender can get zero. This means that
	a return of attacker=1 means that the attacker has lost, and the final remaining piece is the garrison for 
	the square from which the attack was launched.`,
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
			Owner:    args[0],
		}
		defender := army.Army{
			ArmySize: defenderSize,
			Owner:    args[2],
		}
		attacker, defender, err := risk.Run(attacker, defender)

		if err != nil {
			fmt.Printf("ERROR: %v\n, exiting...", err)
		}

		fmt.Printf("Final Standings - %s: %d, %s: %d", attacker.Owner, attacker.ArmySize, defender.Owner, defender.ArmySize)
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
