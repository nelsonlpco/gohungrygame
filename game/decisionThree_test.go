package game

import (
	"testing"

	"github.com/nelsonlpco/gohungrygame/client"
	"github.com/nelsonlpco/gohungrygame/mocks"
)

func Test_ShouldBeCreateAValidDecisionTreeWithThreNodes(t *testing.T) {
	decisionTree := NewDecisionTree("massa", "lasanha", "bolo de chocolate")

	if decisionTree.Data != "massa" {
		t.Errorf("NewDecisionTree.Data is invalid")
	}

	if decisionTree.Left.Data != "lasanha" {
		t.Errorf("NewDecisionTree.Left.Data is invalid")
	}

	if decisionTree.Right.Data != "bolo de chocolate" {
		t.Errorf("NewDecisionTree.Right.Data is invalid")
	}
}

func Test_ShouldBeReturnFalseWhenNodeNotHasChildren(t *testing.T) {
	decisionTree := NewDecisionTree("massa", "lasanha", "bolo de chocolate")

	if decisionTree.HasChildren() == false {
		t.Errorf("NewDecisionTree.HasChildren false, want true")
	}

	if decisionTree.Left.HasChildren() == true {
		t.Errorf("NewDecisionTree.Left.HasChildren true, want false")
	}
}

func Test_ShouldBeAddANewDecicionNode(t *testing.T) {
	decisionTree := NewDecisionTree("massa", "lasanha", "bolo de chocolate")
	node := decisionTree.Right

	node.AddChildren("sorvete", "gelado")

	if node.Data != "gelado" {
		t.Errorf("decisionTree.Right.Data  want gelado")
	}

	if node.Right.Data != "bolo de chocolate" {
		t.Errorf("decisionTree.Right.Right.Data  want bolo de chocolate")
	}

	if node.Left.Data != "sorvete" {
		t.Errorf("decisionTree.Right.Right.Data  want sorvete")
	}
}

func Test_ShouldBeRunTheYesDecisionTree(t *testing.T) {
	cmdClientMock := mocks.BuildCmdClientMock(client.Yes)
	decisionTree := NewDecisionTree("massa", "lasanha", "bolo de chocolate")
	decisionTree.Left.AddChildren("Macarrão", "Comprido")
	decisionTree.Left.Left.AddChildren("Bife", "Carne")

	lastChildren, response := decisionTree.RunCasesWithCli(decisionTree, cmdClientMock)

	if response != client.Yes || lastChildren.Data != "Bife" {
		t.Errorf("lastChildren.data expect Bife, when %s and response expect y when %s", lastChildren.Data, response)
	}

}

func Test_ShouldBeRunTheNoDecisionTree(t *testing.T) {
	cmdClientMock := mocks.BuildCmdClientMock(client.No)
	decisionTree := NewDecisionTree("massa", "lasanha", "bolo de chocolate")
	decisionTree.Right.AddChildren("Sorvete", "Gelado")
	decisionTree.Right.Right.AddChildren("Chocolate", "Tem cacau")

	lastChildren, response := decisionTree.RunCasesWithCli(decisionTree, cmdClientMock)

	if response != client.No || lastChildren.Data != "bolo de chocolate" {
		t.Errorf("lastChildren.data expect Bife, when %s and response expect %s when %s", lastChildren.Data, client.No, response)
	}
}

func Test_ShouldBeQuitOfTheDecisionTreeWhenUserInputQ(t *testing.T) {
	cmdClientMock := mocks.BuildCmdClientMock(client.Quit)
	decisionTree := NewDecisionTree("massa", "lasanha", "bolo de chocolate")
	decisionTree.Right.AddChildren("Sorvete", "Gelado")
	decisionTree.Right.Right.AddChildren("Chocolate", "Tem cacau")

	lastChildren, response := decisionTree.RunCasesWithCli(decisionTree, cmdClientMock)

	if response != client.Quit || lastChildren.Data != "massa" {
		t.Errorf("lastChildren.data expect Bife, when %s and response expect %s when %s", lastChildren.Data, client.No, response)
	}
}
