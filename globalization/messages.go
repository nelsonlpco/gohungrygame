package globalization

var ptBr = map[string]string{
	"ThinkOfADishThatYouLike":                 "Pense em um prato que gosta.",
	"TheDishYouThoughtAboutIs":                "O prato que você pensou é %s ?",
	"IWonAgain":                               "Acertei de novo!",
	"WhatIsTheNameOfTheFood":                  "Qual é o nome da comida?",
	"TheNewFoodIsButTheOtherCurrentFoodIsNot": "%s é ____ que %s não é?",
	"PressToQuit":                             "Pressione [%s] para sair.",
	"YesNo":                                   "[%s | %s]: ",
}

func GetMessage(message string) string {
	return ptBr[message]
}
