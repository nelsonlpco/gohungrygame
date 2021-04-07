package main

import (
	"github.com/nelsonlpco/gohungrygame/client"
	"github.com/nelsonlpco/gohungrygame/cmd"
	"github.com/nelsonlpco/gohungrygame/game"
)

func main() {
	var response string

	decisionTree := game.NewDecisionTree("massa", "lasanha", "bolo de chocolate")
	cmdClient := cmd.NewCmdClient()

	for response != client.Quit {
		cmdClient.ShowStartMessage()
		lastChildren, response := decisionTree.RunCasesWithCli(decisionTree, cmdClient)

		if response == client.Quit {
			break
		}

		if response == client.Yes {
			cmdClient.ShowWinnerMessage()
		} else {
			newFood, quality := cmdClient.NewFood(lastChildren.Data)
			lastChildren.AddChildren(newFood, quality)
		}
	}
}
