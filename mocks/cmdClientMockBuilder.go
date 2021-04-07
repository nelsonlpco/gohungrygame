package mocks

import "github.com/nelsonlpco/gohungrygame/client"

type CmdClientMock struct {
	response string
}

func BuildCmdClientMock(response string) client.Client {
	client := new(CmdClientMock)
	client.response = response

	return client
}

func (c *CmdClientMock) ShowStartMessage() {
}

func (c *CmdClientMock) MakeQuestionWithResponse(food string) string {
	return c.response
}

func (c *CmdClientMock) NewFood(currentFood string) (newFood, quality string) {
	return newFood, quality
}

func (c *CmdClientMock) ShowWinnerMessage() {
}
