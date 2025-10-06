package ui

import (
	"encoding/json"
	"log"

	"github.com/adrianfuro/GoCodePilot/config"
	"github.com/adrianfuro/GoCodePilot/internal/openai"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Helper: render a placeholder tree when loading.
func RenderLoadingTreeView() *tview.TreeView {
	root := tview.NewTreeNode("Loading assistants...").SetColor(tcell.ColorYellow)
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	return tree
}

func StartTUI() {
	app := tview.NewApplication()

	// Initial tree shows "loading"
	tree := RenderLoadingTreeView()
	chat := CreateChatView()
	chat_input := ChatTextInput(chat)

	mainPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(chat, 0, 4, false).
		AddItem(chat_input, 3, 0, false)
	layout := tview.NewFlex().
		AddItem(tree, 30, 1, true).
		AddItem(mainPane, 0, 4, false)

	// Keyboard focus cycling
	focusables := []tview.Primitive{tree, chat, chat_input}
	focusIdx := 0

	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			node := tree.GetCurrentNode()
			asstID, ok := node.GetReference().(string)
			if !ok || asstID == "" {
				return nil
			}
			previousRoot := layout
			promptForTitle(app, previousRoot, func(title string) {
				if title != "" {
					configPath := config.DefaultConfigPath()
					envVars := config.LoadEnvVars()
					client := openai.NewClient(envVars.OpenAIKey)
					err := config.AddThreadToAssistant(configPath, asstID, title, client)
					if err != nil {
						log.Printf("Error adding thread: %v", err)
						return
					}
					assistants, err := config.LoadAssistantConfigs(configPath)
					if err != nil {
						log.Printf("Error reloading assistants config: %v", err)
						return
					}
					app.QueueUpdateDraw(func() {
						newTree := RenderTreeView(assistants)
						layout.RemoveItem(tree)
						layout.AddItem(newTree, 30, 1, true)
						tree = newTree
						app.SetFocus(tree)
					})
				}
			})
			return nil
		}
		return event
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			focusIdx = (focusIdx + 1) % len(focusables)
			app.SetFocus(focusables[focusIdx])
			return nil
		case tcell.KeyBacktab:
			focusIdx = (focusIdx - 1 + len(focusables)) % len(focusables)
			app.SetFocus(focusables[focusIdx])
			return nil
		}
		return event
	})

	app.EnableMouse(true)
	app.SetRoot(layout, true)

	// --- FAST ASYNC LOADING: assistants will be loaded in background and UI updated ---
	go func() {
		configPath := config.DefaultConfigPath()
		envVars := config.LoadEnvVars()
		client := openai.NewClient(envVars.OpenAIKey)
		apiJSON, err := client.ListAssistantOpenAI()
		if err != nil {
			log.Printf("Failed to fetch assistants JSON from OpenAI: %v", err)
			return
		}
		var apiAssistants []config.AssistantConfig
		if err := json.Unmarshal([]byte(apiJSON), &apiAssistants); err != nil {
			log.Printf("Failed to unmarshal assistants JSON: %v", err)
			return
		}
		if err := config.SyncAssistantsFromAPI(configPath, apiAssistants); err != nil {
			log.Printf("Failed to sync assistants config: %v", err)
			return
		}
		assistants, err := config.LoadAssistantConfigs(configPath)
		if err != nil {
			log.Printf("Failed to load assistants config: %v", err)
			return
		}

		app.QueueUpdateDraw(func() {
			newTree := RenderTreeView(assistants)
			layout.RemoveItem(tree)
			layout.AddItem(newTree, 30, 1, true)
			tree = newTree
			app.SetFocus(tree)
		})
	}()
	// --- END ASYNC LOADING ---

	if err := app.Run(); err != nil {
		log.Fatalf("TUI app failed: %v", err)
	}
}
