package cmd

import (
	"fmt"

	"github.com/nelsonlpco/gohungrygame/client"
	glo "github.com/nelsonlpco/gohungrygame/globalization"
)

type CmdClient struct{}

func NewCmdClient() *CmdClient {
	return new(CmdClient)
}

func (c CmdClient) ShowStartMessage() {
	fmt.Printf(glo.GetMessage("PressToQuit")+"\n\n", client.Quit)
	fmt.Println(glo.GetMessage("ThinkOfADishThatYouLike") + "\n")
}

func (c CmdClient) MakeQuestionWithResponse(food string) string {
	var response string
	yesNo := fmt.Sprintf(glo.GetMessage("YesNo"), client.Yes, client.No)

	for response != client.Yes && response != client.No && response != client.Quit {
		fmt.Printf(glo.GetMessage("TheDishYouThoughtAboutIs")+yesNo, food)
		fmt.Scanln(&response)
	}

	return response
}

func (c CmdClient) NewFood(currentFood string) (newFood, quality string) {
	for newFood == "" {
		fmt.Print(glo.GetMessage("WhatIsTheNameOfTheFood"))
		fmt.Scanln(&newFood)
	}

	for quality == "" {
		fmt.Printf(glo.GetMessage("TheNewFoodIsButTheOtherCurrentFoodIsNot"), newFood, currentFood)
		fmt.Scanln(&quality)
	}

	return newFood, quality
}

func (c CmdClient) ShowWinnerMessage() {
	fmt.Println("\n\n\n-------------------------")
	fmt.Println(glo.GetMessage("IWonAgain"))
	fmt.Print("-------------------------\n\n\n")
}
