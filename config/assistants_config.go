package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrianfuro/GoCodePilot/internal/openai"
)

type Tool struct {
	Type string `json:"type"`
}

type Message struct {
	ID      string `json:"id"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Thread struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Messages []Message `json:"messages"`
}

type AssistantConfig struct {
	ID           string   `json:"id"`
	Object       string   `json:"object"`
	Name         string   `json:"name"`
	Model        string   `json:"model"`
	Instructions string   `json:"instructions"`
	Tools        []Tool   `json:"tools"`
	Threads      []Thread `json:"threads"` // local only!
}

type APIAssistant struct {
	ID           string `json:"id"`
	Object       string `json:"object"`
	Name         string `json:"name"`
	Model        string `json:"model"`
	Instructions string `json:"instructions"`
	Tools        []Tool `json:"tools"`
}

func LoadAssistantConfigs(path string) ([]AssistantConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []AssistantConfig{}, nil // no config yet
		}
		return nil, err
	}
	defer f.Close()
	var configs []AssistantConfig
	err = json.NewDecoder(f).Decode(&configs)
	return configs, err
}

func SaveAssistantConfigs(path string, configs []AssistantConfig) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(configs)
}

func DefaultConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		// fallback to home or local directory
		return "assistants.json"
	}
	appDir := filepath.Join(configDir, "gocodepilot")
	os.MkdirAll(appDir, 0755) // ensure directory exists
	return filepath.Join(appDir, "assistants.json")
}

func SyncAssistantsFromAPI(path string, apiAssistants []AssistantConfig) error {
	localConfigs, _ := LoadAssistantConfigs(path)

	// Build API IDs map for quick lookup
	apiIDs := make(map[string]struct{}, len(apiAssistants))
	for _, a := range apiAssistants {
		apiIDs[a.ID] = struct{}{}
	}

	// Build a new slice only with assistants present in the API
	var syncedConfigs []AssistantConfig
	configMap := make(map[string]*AssistantConfig)
	for i := range localConfigs {
		configMap[localConfigs[i].ID] = &localConfigs[i]
	}

	// Update/add all API assistants
	for _, a := range apiAssistants {
		if existing, found := configMap[a.ID]; found {
			// Update fields, keep threads
			existing.Object = a.Object
			existing.Name = a.Name
			existing.Model = a.Model
			existing.Instructions = a.Instructions
			existing.Tools = a.Tools
			syncedConfigs = append(syncedConfigs, *existing)
		} else {
			// New assistant from API
			syncedConfigs = append(syncedConfigs, AssistantConfig{
				ID:           a.ID,
				Object:       a.Object,
				Name:         a.Name,
				Model:        a.Model,
				Instructions: a.Instructions,
				Tools:        a.Tools,
				Threads:      []Thread{},
			})
		}
	}

	// Save only the assistants that exist in the API
	return SaveAssistantConfigs(path, syncedConfigs)
}

func AddThreadToAssistant(configPath string, assistantID, title string, client *openai.Client) error {
	threadResp, err := client.CreateThread()
	if err != nil {
		return err
	}
	threadID := threadResp.ID // Use the ID returned by OpenAI

	assistants, err := LoadAssistantConfigs(configPath)
	if err != nil {
		return err
	}

	updated := false
	for i, asst := range assistants {
		if asst.ID == assistantID {
			assistants[i].Threads = append(assistants[i].Threads, Thread{
				ID:    threadID,
				Title: title,
			})
			updated = true
			break
		}
	}
	if !updated {
		return fmt.Errorf("assistant not found: %s", assistantID)
	}
	return SaveAssistantConfigs(configPath, assistants)
}
