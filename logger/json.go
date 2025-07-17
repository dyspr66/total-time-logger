package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"ttl/data"
)

func (t *TotalTimeLogger) ReadFromJson(ttl *TotalTimeLogger) error {
	// Read file
	b, err := os.ReadFile("record.json")
	if err != nil {
		return fmt.Errorf("reading file: %w", err)
	}

	if len(b) == 0 {
		return nil
	}

	// Use file data in program
	var acts []data.Activity
	err = json.Unmarshal(b, &acts)
	if err != nil {
		return fmt.Errorf("unmarshaling from json: %w", err)
	}

	var a data.Activities
	for _, v := range acts {
		a = append(a, &v)
	}

	ttl.Activities = a

	return nil
}

func (t *TotalTimeLogger) SaveToJSON() error {
	// Structute the program data
	var acts []data.Activity
	for _, v := range t.Activities {
		acts = append(acts, *v)
	}

	// Write program data into file
	jsonData, err := json.Marshal(acts)
	if err != nil {
		return fmt.Errorf("marshaling to json: %w", err)
	}

	err = os.WriteFile("record.json", jsonData, 0644)
	if err != nil {
		return fmt.Errorf("writing data to json file: %w", err)
	}

	return nil
}
