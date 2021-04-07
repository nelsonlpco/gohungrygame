package client

type Client interface {
	ShowStartMessage()
	MakeQuestionWithResponse(food string) string
	NewFood(currentFood string) (newFood, quality string)
	ShowWinnerMessage()
}

const (
	Yes  = "s"
	No   = "n"
	Quit = "q"
)
