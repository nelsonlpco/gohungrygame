package game

import "github.com/nelsonlpco/gohungrygame/client"

type Node struct {
	Left  *Node
	Right *Node
	Data  string
}

func NewDecisionTree(quality, leftFood, rightFood string) *Node {
	root := new(Node)
	root.Data = quality

	root.Left = &Node{Data: leftFood}
	root.Right = &Node{Data: rightFood}

	return root
}

func (n *Node) HasChildren() bool {
	return n.Left != nil && n.Right != nil
}

func (n *Node) AddChildren(food, quality string) {
	n.Right = &Node{Data: n.Data}
	n.Data = quality
	n.Left = &Node{Data: food}
}

func (n Node) RunCasesWithCli(decisionTree *Node, cli client.Client) (lastChildren *Node, response string) {
	lastChildren = decisionTree
	response = cli.MakeQuestionWithResponse(lastChildren.Data)

	if response == client.Quit {
		return lastChildren, response
	}

	if lastChildren.HasChildren() {
		if response == client.Yes {
			return lastChildren.RunCasesWithCli(lastChildren.Left, cli)
		} else if response == client.No {
			return lastChildren.RunCasesWithCli(lastChildren.Right, cli)
		}
	}

	return lastChildren, response
}
