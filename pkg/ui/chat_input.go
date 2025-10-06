package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ChatTextInput(chat *tview.TextView) *tview.InputField {
	chat_input := tview.NewInputField()

	chat_input.SetBorder(true)
	chat_input.SetLabel("Message:").
		SetFieldBackgroundColor(tview.Styles.ContrastBackgroundColor).
		SetFieldTextColor(tview.Styles.PrimaryTextColor).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				msg := chat_input.GetText()
				if msg != "" {
					// Inside SetDoneFunc:
					formatted := "[green]User >> [white]" + msg
					currentText := chat.GetText(true)
					if currentText != "" {
						chat.SetText(currentText + "\n" + formatted)
					} else {
						chat.SetText(formatted)
					}

					chat_input.SetText("")
				}
			}
		})

	return chat_input
}
