package ui

import (
	"github.com/adrianfuro/GoCodePilot/config"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderTreeView(assistants []config.AssistantConfig) *tview.TreeView {
	root := tview.NewTreeNode("Assistants").SetColor(tcell.ColorRed).SetExpanded(true)

	for _, asst := range assistants {
		asstNode := tview.NewTreeNode(asst.Name).SetColor(tcell.ColorGreen).SetReference(asst.ID)
		for _, thread := range asst.Threads {
			threadNode := tview.NewTreeNode(thread.Title).SetReference(thread.ID)
			asstNode.AddChild(threadNode)
		}
		root.AddChild(asstNode)
	}

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	tree.SetBorder(true).
		SetTitle("Chat History")

	return tree
}

func promptForTitle(app *tview.Application, previousRoot tview.Primitive, onSubmit func(title string)) {
	input := tview.NewInputField().
		SetLabel("Thread Title: ")

	input.SetFieldWidth(40).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				title := input.GetText()
				app.SetRoot(previousRoot, true) // Restore main layout
				onSubmit(title)
			}
			if key == tcell.KeyEsc {
				app.SetRoot(previousRoot, true) // Cancel, just restore
			}
		})

	modal := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(input, 3, 1, true).
		AddItem(nil, 0, 1, false)

	app.SetRoot(modal, true)
	app.SetFocus(input)
}
