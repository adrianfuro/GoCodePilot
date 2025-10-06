package ui

import (
	"github.com/rivo/tview"
)

func CreateChatView() *tview.TextView {
	chat_ui := tview.NewTextView()

	chat_ui.SetBorder(true).SetTitle("Chat")
	chat_ui.
		SetDynamicColors(true)

	return chat_ui
}
