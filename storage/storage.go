package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Task struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

type Data struct {
	Groups map[string][]Task `json:"groups"`
}

var dataFile string

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dataFile = filepath.Join(homeDir, ".combo", "data", "tasks.json")
}

// Load 加载数据
func Load() (*Data, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return &Data{Groups: map[string][]Task{"default": {}}}, nil
	}
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}
	var d Data
	if err := json.Unmarshal(file, &d); err != nil {
		return nil, err
	}
	if d.Groups == nil {
		d.Groups = make(map[string][]Task)
	}
	return &d, nil
}

// Save 保存数据
func Save(d *Data) error {
	err := os.MkdirAll(filepath.Dir(dataFile), os.ModePerm)
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, bytes, 0644)
}
